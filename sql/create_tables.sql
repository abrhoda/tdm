CREATE TYPE licenses AS ENUM('OGL', 'ORC');
CREATE TYPE rarities AS ENUM('common', 'uncommon', 'rare', 'unique');
CREATE TYPE attributes AS ENUM('strength', 'dexterity', 'consitution', 'intelligence', 'wisdom', 'charisma');
CREATE TYPE sizes AS ENUM('small', 'medium', 'large', 'huge', 'gargantuan')
CREATE TYPE visions AS ENUM('normal', 'low light vision', 'dark vision')

CREATE TABLE trait IF NOT EXISTS (
  trait_id int PRIMARY KEY GENERATED ALWAYS AS INDENTITY,
  value text NOT NULL,
  -- TODO ADD COLUMN description NOT NULL,
);

CREATE TABLE ancestries_traits IF NOT EXISTS (
  ancestry_id int REFERENCES ancestry ON DELETE CASCADE,
  trait_id int REFERENCES trait ON DELETE CASCADE,
  PRIMARY KEY (ancestry_id, trait_id),
);

CREATE TABLE language IF NOT EXISTS (
  language_id int PRIMARY KEY GENERATED ALWAYS AS INDENTITY,
  value text NOT NULL,
);

CREATE TABLE ancestries_languages IF NOT EXISTS (
  ancestry_id int REFERENCES ancestry ON DELETE CASCADE,
  language_id int REFERENCES language ON DELETE CASCADE,
  is_additional_language boolean NOT NULL,
  PRIMARY KEY (ancestry_id, language_id),
);

CREATE TABLE ancestry IF NOT EXISTS (
  ancestry_id int PRIMARY KEY GENERATED ALWAYS AS INDENTITY,
  name text NOT NULL,
  description text NOT NULL,
  game_master_description text,
  title text NOT NULL,
  remaster boolean NOT NULL,
  license licenses NOT NULL,
  rarity rarities NOT NULL,
  rules text, -- turn all rules into a text blob for now
  first_boost attributes NOT NULL,
  free_boot text NOT NULL DEFAULT 'free',
  second_boost attributes,
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

CREATE TABLE ancestry_feature IF NOT EXISTS (
  ancestry_feature_id int PRIMARY KEY GENERATED ALWAYS AS INDENTITY,
  ancestry_id int REFERNECES ancestry ON DELETE CASCADE,
  name text NOT NULL,
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

);

CREATE TABLE ancestry_features_traits IF NOT EXISTS (
  ancestry_feature_id int REFERENCES ancestry_feature ON DELETE CASCADE,
  trait_id int REFERENCES trait ON DELETE CASCADE,
  PRIMARY KEY (ancestry_feature_id, trait_id),
);

CREATE TABLE proficiency IF NOT EXISTS (
  proficiency_id int PRIMARY KEY GENERATED ALWAYS AS INDENTITY,
  name text NOT NULL,
  -- called value and not rank because pgsql uses rank function
  value int CHECK (0 <= value AND value <= 4)
);

-- junction table for this one to one mapping because proficiencies will also map to othe r tables
CREATE TABLE ancestries_proficiencies IF NOT EXISTS (
  ancestry_id int REFERENCES ancestry ON DELETE CASCADE,
  proficiency_id int REFERENCES proficiency ON DELETE CASCADE,
  PRIMARY KEY (ancestry_id, proficiency_id),
);

CREATE TABLE sense IF NOT EXISTS (
  sense_id int PRIMARY KEY GENERATED ALWAYS AS INDENTITY,
  name text NOT NULL,
  acuity int,
  range int,
  elevate_if_has_low_light_vision boolean,
  CONSTRAINT valid_elevate_field CHECK 
    ((name = 'low light vision' AND elevate_if_has_low_light_vision IS NOT NULL) OR 
    (name != 'low light vision' AND elevate_if_has_low_light_vision IS NULL))
);

-- junction table for this one to one mapping because senses will also map to othe r tables
CREATE TABLE ancestry_features_senses IF NOT EXISTS (
  ancestry_id int REFERENCES ancestry ON DELETE CASCADE,
  sense_id int REFERENCES sense ON DELETE CASCADE,
  PRIMARY KEY (ancestry_id, sense_id),
);
