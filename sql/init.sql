CREATE TYPE licenses AS ENUM('ogl', 'orc');
CREATE TYPE rarities AS ENUM('common', 'uncommon', 'rare', 'unique');
CREATE TYPE abilities AS ENUM('strength', 'dexterity', 'constitution', 'intelligence', 'wisdom', 'charisma');
CREATE TYPE sizes AS ENUM('small', 'medium', 'large', 'huge', 'gargantuan');
CREATE TYPE visions AS ENUM('normal', 'low light vision', 'dark vision');
CREATE TYPE boost_types AS ENUM('first', 'second', 'third', 'flaw');
CREATE TYPE periods AS ENUM('day', 'turn', 'round', 'PT10M', 'PT1H', 'PT1M');
CREATE TYPE proficiency_categories AS ENUM('class dc', 'spellcasting dc', 'saving throw', 'defense', 'attack', 'skill');
CREATE TYPE proficiency_ranks AS ENUM('untrained', 'trained', 'expert', 'master', 'legendary');
CREATE TYPE feat_types AS ENUM('ancestry', 'class', 'general', 'skill');

CREATE TABLE IF NOT EXISTS trait (
  id int PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  value text UNIQUE NOT NULL
  -- TODO ADD COLUMN description NOT NULL,
);

CREATE TABLE IF NOT EXISTS language (
  id int PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  value text UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS sense (
  id int PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  name text NOT NULL, -- not unique!
  acuity int,
  range int,
  elevate_if_has_low_light_vision boolean,

  CONSTRAINT valid_elevate_field CHECK 
    ((name = 'low light vision' AND elevate_if_has_low_light_vision IS NOT NULL) OR
    (name != 'low light vision' AND elevate_if_has_low_light_vision IS NULL))
);

CREATE TABLE IF NOT EXISTS boost (
  id int PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  ability abilities NOT NULL
);

CREATE TABLE IF NOT EXISTS proficiency (
  id int PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  name text NOT NULL,
  -- called value and not rank because pgsql uses rank function
  category proficiency_categories NOT NULL,
  proficiency_rank proficiency_ranks NOT NULL
);

CREATE TABLE IF NOT EXISTS ancestry (
  id int PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  name text UNIQUE NOT NULL,
  description text NOT NULL,
  game_master_description text,
  title text NOT NULL,
  remaster boolean NOT NULL,
  license licenses NOT NULL,
  rarity rarities NOT NULL,
  rules text, -- turn all rules into a text blob for now
  additional_language_count int NOT NULL,
  flaw abilities,
  hp int NOT NULL,
  reach int NOT NULL,
  size sizes NOT NULL,
  speed int NOT NULL,
  vision visions NOT NULL,

  CONSTRAINT validate_hp_in_range check (6 <= hp AND hp <= 12 AND hp % 2 = 0)
  -- TODO validate this in the ancestries_boosts table instead.
  --CONSTRAINT validate_flaw_if_second_boost_present CHECK ((second_boost IS NOT NULL AND flaw IS NOT NULL) OR (second_boost IS NULL AND flaw IS NULL))
);

-- junction table for this one to many mapping because boosts will also map to othe r tables
CREATE TABLE IF NOT EXISTS ancestries_boosts (
  ancestry_id int REFERENCES ancestry ON DELETE CASCADE,
  boost_id int REFERENCES boost ON DELETE CASCADE,
  boost_type boost_types NOT NULL,
  PRIMARY KEY (ancestry_id, boost_id)
);

-- junction table for this one to many mapping because traits will also map to othe r tables
CREATE TABLE IF NOT EXISTS ancestries_traits (
  ancestry_id int REFERENCES ancestry ON DELETE CASCADE,
  trait_id int REFERENCES trait ON DELETE CASCADE,
  PRIMARY KEY (ancestry_id, trait_id)
);

-- junction table for this one to many mapping because languages will also map to othe r tables
CREATE TABLE IF NOT EXISTS ancestries_languages (
  ancestry_id int REFERENCES ancestry ON DELETE CASCADE,
  language_id int REFERENCES language ON DELETE CASCADE,
  is_additional_language boolean NOT NULL,
  PRIMARY KEY (ancestry_id, language_id)
);

CREATE TABLE IF NOT EXISTS ancestry_property (
  ancestry_property_id int PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  ancestry_id int REFERENCES ancestry ON DELETE CASCADE,
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

  CONSTRAINT validate_grants_lang_count_if_present check ( grants_languages_count >= 0),
  CONSTRAINT validate_level check ( level >= 0)
);

-- junction table for this one to many mapping because languages will also map to othe r tables
CREATE TABLE IF NOT EXISTS ancestry_properties_languages (
  ancestry_properties_id int REFERENCES ancestry_property ON DELETE CASCADE,
  language_id int REFERENCES language ON DELETE CASCADE,
  PRIMARY KEY (ancestry_properties_id, language_id)
);

CREATE TABLE IF NOT EXISTS ancestry_properties_traits (
  ancestry_property_id int REFERENCES ancestry_property ON DELETE CASCADE,
  trait_id int REFERENCES trait ON DELETE CASCADE,
  PRIMARY KEY (ancestry_property_id, trait_id)
);

-- junction table for this one to many mapping because senses will also map to othe r tables
CREATE TABLE IF NOT EXISTS ancestry_properties_senses (
  ancestry_property_id int REFERENCES ancestry_property ON DELETE CASCADE,
  sense_id int REFERENCES sense ON DELETE CASCADE,
  PRIMARY KEY (ancestry_property_id, sense_id)
);

CREATE TABLE IF NOT EXISTS background (
  id int PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  name text UNIQUE NOT NULL,
  description text NOT NULL,
  game_master_description text,
  title text NOT NULL,
  remaster boolean NOT NULL,
  license licenses NOT NULL,
  rarity rarities NOT NULL,
  rules text
);

-- junction table for this one one to many mapping because skills will also map to other tables
CREATE TABLE IF NOT EXISTS backgrounds_proficiencies (
  background_id int REFERENCES background ON DELETE CASCADE,
  proficiency_id int REFERENCES proficiency ON DELETE CASCADE,
  PRIMARY KEY (background_id, proficiency_id)
);

-- create general feats
CREATE TABLE IF NOT EXISTS general_feat (
  id int PRIMARY KEY GENERATED ALWAYS AS IDENTITY, 
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
  PRIMARY KEY (background_id, general_feat_id)
);

CREATE TABLE IF NOT EXISTS backgrounds_traits (
  background_id int REFERENCES background ON DELETE CASCADE,
  trait_id int REFERENCES trait ON DELETE CASCADE,
  PRIMARY KEY (background_id, trait_id)
);

CREATE TABLE IF NOT EXISTS general_feats_traits (
  general_feat_id int REFERENCES general_feat ON DELETE CASCADE,
  trait_id int REFERENCES trait ON DELETE CASCADE,
  PRIMARY KEY (general_feat_id, trait_id)
);

-- create prerequisite table
CREATE TABLE IF NOT EXISTS prerequisite (
  id int PRIMARY KEY GENERATED ALWAYS AS IDENTITY, 
  value text NOT NULL
);

CREATE TABLE IF NOT EXISTS general_feats_prerequisites (
  general_feat_id int REFERENCES general_feat ON DELETE CASCADE,
  prerequisite_id int REFERENCES prerequisite ON DELETE CASCADE,
  PRIMARY KEY (general_feat_id, prerequisite_id)
);

CREATE TABLE IF NOT EXISTS feat_level (
  id int PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  level int NOT NULL,
  feat_type feat_types NOT NULL
);

CREATE TABLE IF NOT EXISTS key_ability (
  id int PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  ability abilities NOT NULL
);

CREATE TABLE IF NOT EXISTS class (
  id int PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  name text UNIQUE NOT NULL,
  description text NOT NULL,
  game_master_description text,
  title text NOT NULL,
  remaster boolean NOT NULL,
  license licenses NOT NULL,
  rarity rarities NOT NULL,
  rules text,
  hit_points int NOT NULL,
  perception proficiency_ranks NOT NULL,
  spellcasting proficiency_ranks NOT NULL,
  additional_trained_skills int NOT NULL,

  CONSTRAINT validate_hit_point_range check (6 <= hit_points OR hit_points <= 12),
  CONSTRAINT validate_additional_trained_skills_count check (additional_trained_skills >= 0)
);

CREATE TABLE IF NOT EXISTS classes_traits (
  class_id int REFERENCES class ON DELETE CASCADE,
  trait_id int REFERENCES trait ON DELETE CASCADE,
  PRIMARY KEY (class_id, trait_id)
);

CREATE TABLE IF NOT EXISTS classes_feat_levels (
  class_id int REFERENCES class ON DELETE CASCADE,
  feat_level_id int REFERENCES feat_level ON DELETE CASCADE,
  PRIMARY KEY (class_id, feat_level_id)
);

CREATE TABLE IF NOT EXISTS classes_key_abilities (
  class_id int REFERENCES class ON DELETE CASCADE,
  key_ability_id int REFERENCES key_ability ON DELETE CASCADE,
  PRIMARY KEY (class_id, key_ability_id)
);

-- junction table for this one one to many mapping because proficiencies will also map to other tables
CREATE TABLE IF NOT EXISTS classes_proficiencies (
  class_id int REFERENCES class ON DELETE CASCADE,
  proficiency_id int REFERENCES proficiency ON DELETE CASCADE,
  PRIMARY KEY (class_id, proficiency_id)
);


CREATE TABLE IF NOT EXISTS class_property (
  id int PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  class_id int REFERENCES class ON DELETE CASCADE,
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
  level int NOT NULL
);

CREATE TABLE IF NOT EXISTS tag (
  tag_id int PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  value text UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS class_properties_tags (
  class_property_id int REFERENCES class_property ON DELETE CASCADE,
  tag_id int REFERENCES tag ON DELETE CASCADE,
  PRIMARY KEY (class_property_id, tag_id)
);

CREATE TABLE IF NOT EXISTS class_properties_traits (
  class_property_id int REFERENCES class_property ON DELETE CASCADE,
  trait_id int REFERENCES trait ON DELETE CASCADE,
  PRIMARY KEY (class_property_id, trait_id)
);

CREATE TABLE IF NOT EXISTS class_properties_key_abilities (
  class_property_id int REFERENCES class_property ON DELETE CASCADE,
  key_ability_id int REFERENCES key_ability ON DELETE CASCADE,
  PRIMARY KEY (class_property_id, key_ability_id)
);

CREATE TABLE IF NOT EXISTS class_properties_proficiencies (
  class_property_id int REFERENCES class_property ON DELETE CASCADE,
  proficiency_id int REFERENCES proficiency ON DELETE CASCADE,
  PRIMARY KEY (class_property_id, proficiency_id)
);

-- todo add ancestry, bonus, class, and skill feats table
-- todo add feat effects table
-- todo junction table for ancestry_feats_feat_effects, class_feats_feat_effects and, skill_feats_feat_effects 

