ALTER TABLE variants DROP INDEX flag_key_key;
ALTER TABLE variants ADD UNIQUE(`key`);
