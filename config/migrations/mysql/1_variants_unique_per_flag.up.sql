
ALTER TABLE variants DROP INDEX `key`;
ALTER TABLE variants ADD CONSTRAINT variants_flag_key_key_key UNIQUE(flag_key, `key`);