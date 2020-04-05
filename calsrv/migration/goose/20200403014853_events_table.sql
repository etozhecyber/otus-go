-- +goose Up
CREATE TABLE IF NOT EXISTS events (
  id UUID DEFAULT uuid_generate_v4() PRIMARY KEY  ,
  Owner_id integer NOT NULL,
  Title text NOT NULL,
  Text  text,
  StartTime timestamp without time zone NOT NULL,
  EndTime timestamp without time zone NOT NULL,
  foreign key(Owner_id) references users(id)
  );

CREATE INDEX owner_idx ON events (Owner_id);
CREATE INDEX StartTime_idx ON events (StartTime);

-- +goose Down
DROP TABLE events;
DROP INDEX owner_idx;
DROP INDEX StartTime_idx;
