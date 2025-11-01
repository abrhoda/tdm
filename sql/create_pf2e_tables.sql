CREATE TYPE licenses AS ENUM('OGL', 'ORC');
CREATE TYPE rarities AS ENUM('common', 'uncommon', 'rare', 'unique');
CREATE TYPE attributes AS ENUM('any', 'strength', 'dexterity', 'consitution', 'intelligence', 'wisdom', 'charisma');
CREATE TYPE sizes AS ENUM('small', 'medium', 'large', 'huge', 'gargantuan');
CREATE TYPE visions AS ENUM('normal', 'low light vision', 'dark vision');
CREATE TYPE boost_types AS ENUM('first', 'second', 'third', 'flaw');
CREATE TYPE periods AS ENUM('day', 'round', 'PT10M', 'PT1H', 'PT1M', );

CREATE TABLE IF NOT EXISTS trait (
  trait_id int PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  value text UNIQUE NOT NULL,
  -- TODO ADD COLUMN description NOT NULL,
);

CREATE TABLE IF NOT EXISTS language (
  language_id int PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  value text UNIQUE NOT NULL,
);

CREATE TABLE IF NOT EXISTS sense (
  sense_id int PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  name text NOT NULL, -- not unique!
  acuity int,
  range int,
  elevate_if_has_low_light_vision boolean,

  CONSTRAINT valid_elevate_field CHECK 
    ((name = 'low light vision' AND elevate_if_has_low_light_vision IS NOT NULL) OR 
    (name != 'low light vision' AND elevate_if_has_low_light_vision IS NULL))
);

CREATE TABLE IF NOT EXISTS boost (
  boost_id int PRIMARY KEY GENERATED ALWAYS AS IDENTITY
  attribute attributes NOT NULL,
);

CREATE TABLE IF NOT EXISTS proficiency (
  proficiency_id int PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  name text NOT NULL,
  -- called value and not rank because pgsql uses rank function
  value int CHECK (0 <= value AND value <= 4)
);

CREATE TABLE IF NOT EXISTS ancestry (
  ancestry_id int PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  name text UNIQUE NOT NULL,
  description text NOT NULL,
  game_master_description text,
  title text NOT NULL,
  remaster boolean NOT NULL,
  license licenses NOT NULL,
  rarity rarities NOT NULL,
  rules text, -- turn all rules into a text blob for now
  additional_language_count int NOT NULL,
  flaw attributes,
  hp int NOT NULL,
  reach int NOT NULL,
  size sizes NOT NULL,
  speed int NOT NULL,
  vision visions NOT NULL,

  CONSTRAINT validate_hp_in_range check (6 <= hp AND hp <= 12 AND hp % 2 == 0)
  CONSTRAINT validate_flaw_if_second_boost_present CHECK ((second_boost IS NOT NULL AND flaw IS NOT NULL) OR (second_boost IS NULL AND flaw IS NULL))
);

-- junction table for this one to many mapping because boosts will also map to othe r tables
CREATE TABLE IF NOT EXISTS ancestries_boosts (
  ancestry_id int REFERENCES ancestry ON DELETE CASCADE,
  boost_id int REFERENCES boost ON DELETE CASCADE,
  boost_type boost_types NOT NULL,
  PRIMARY KEY (ancestry_id, boost_id),
);

-- junction table for this one to many mapping because traits will also map to othe r tables
CREATE TABLE IF NOT EXISTS ancestries_traits (
  ancestry_id int REFERENCES ancestry ON DELETE CASCADE,
  trait_id int REFERENCES trait ON DELETE CASCADE,
  PRIMARY KEY (ancestry_id, trait_id),
);

-- junction table for this one to many mapping because languages will also map to othe r tables
CREATE TABLE IF NOT EXISTS ancestries_languages (
  ancestry_id int REFERENCES ancestry ON DELETE CASCADE,
  language_id int REFERENCES language ON DELETE CASCADE,
  is_additional_language boolean NOT NULL,
  PRIMARY KEY (ancestry_id, language_id),
);

CREATE TABLE IF NOT EXISTS ancestry_property (
  ancestry_property_id int PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  ancestry_id int REFERNECES ancestry ON DELETE CASCADE,
  name text UNIQUE NOT NULL,
  description text NOT NULL,
  game_master_description text,
  title text NOT NULL,
  remaster boolean NOT NULL,
  license licenses NOT NULL,
  rarity rarities NOT NULL,
  rules text, -- turn all rules into a text blob for now

  -- might not need these 3 columns.
  action_type text NOT NULL,
  actions int,
  category text NOT NULL,
  ---

  level int NOT NULL, -- might always be 0
  grants_languages_count int DEFAULT 0,

  CONSTRAINT validate_grants_lang_count_if_present check ( grants_language_count >= 0)
  CONSTRAINT validate_level check ( level >= 0)
);

-- junction table for this one to many mapping because languages will also map to othe r tables
CREATE TABLE IF NOT EXISTS ancestry_properties_languages (
  ancestry_properties_id int REFERENCES ancestry_property ON DELETE CASCADE,
  language_id int REFERENCES language ON DELETE CASCADE,
  PRIMARY KEY (ancestry_properties_id, language_id),
);

CREATE TABLE IF NOT EXISTS ancestry_properties_traits (
  ancestry_property_id int REFERENCES ancestry_property ON DELETE CASCADE,
  trait_id int REFERENCES trait ON DELETE CASCADE,
  PRIMARY KEY (ancestry_property_id, trait_id),
);

-- junction table for this one to many mapping because senses will also map to othe r tables
CREATE TABLE IF NOT EXISTS ancestry_properties_senses (
  ancestry_property_id int REFERENCES ancestry_property ON DELETE CASCADE,
  sense_id int REFERENCES sense ON DELETE CASCADE,
  PRIMARY KEY (ancestry_property_id, sense_id),
);

CREATE TABLE IF NOT EXISTS background (
  background_id int PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  name text UNIQUE NOT NULL,
  description text NOT NULL,
  game_master_description text,
  title text NOT NULL,
  remaster boolean NOT NULL,
  license licenses NOT NULL,
  rarity rarities NOT NULL,
  rules text,
);

CREATE TABLE IF NOT EXISTS skill (
  skill_id int PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  name text UNIQUE NOT NULL,
);

-- junction table for this one one to many mapping because skills will also map to other tables
CREATE TABLE IF NOT EXISTS backgrounds_skills (
  background_id int REFERENCES background ON DELETE CASCADE,
  skill_id int REFERENCES skill ON DELETE CASCADE,
  PRIMARY KEY (background_id, skill_id),
);

-- create general feats
CREATE TABLE IF NOT EXISTS general_feat (
  general_feat_id int PRIMARY KEY GENERATED ALWAYS AS IDENTITY, 
  name text UNIQUE NOT NULL,
  description text NOT NULL,
  game_master_description text,
  title text NOT NULL,
  remaster boolean NOT NULL,
  license licenses NOT NULL,
  rarity rarities NOT NULL,
  rules text,
  action_type text NOT NULL,
  actions int,
  category text NOT NULL,
  level int NOT NULL,
  max_takable int NOT NULL DEFAULT 1,
  frequency_max int,
  frequency_period periods,

  CONSTRAINT validate_level check ( level >= 0),
  CONSTRAINT validate_max_takable check ( max_takable > 0 ),
  CONSTRAINT validate_frequency_max check ( frequency_max IS NULL OR frequency_max > 0 ),
  CONSTRAINT validate_frequency_period_if_has_frequency_max check (
    (frequency_max IS NULL and frequency_period IS NULL) OR
    (frequency_max IS NOT NULL AND frequency_period IS NOT NULL)
  )
);

-- a few backgrounds grant more than 1 feat. 
CREATE TABLE IF NOT EXISTS backgrounds_general_feats (
  background_id int REFERENCES background ON DELETE CASCADE,
  general_feat_id int REFERENCES general_feat ON DELETE CASCADE,
  PRIMARY KEY (background_id, general_feat_id),
);

CREATE TABLE IF NOT EXISTS backgrounds_traits (
  background_id int REFERENCES background ON DELETE CASCADE,
  trait_id int REFERENCES trait ON DELETE CASCADE,
  PRIMARY KEY (background_id, trait_id),
);

CREATE TABLE IF NOT EXISTS general_feats_traits (
  general_feat_id int REFERENCES general_feat ON DELETE CASCADE,
  trait_id int REFERENCES trait ON DELETE CASCADE,
  PRIMARY KEY (general_feat_id, trait_id),
);

-- create prerequisite table
CREATE TABLE IF NOT EXISTS prerequisite (
  prerequisite_id int PRIMARY KEY GENERATED ALWAYS AS IDENTITY, 
  value text NOT NULL,
);

CREATE TABLE IF NOT EXISTS general_feats_prerequisites (
  general_feat_id int REFERENCES general_feat ON DELETE CASCADE,
  prerequisite_id int REFERENCES prerequisite ON DELETE CASCADE,
  PRIMARY KEY (general_feat_id, prerequisite_id),
);


