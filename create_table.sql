-- Create table
CREATE TABLE IF NOT EXISTS vehicle_locations (
    id BIGSERIAL PRIMARY KEY,
    vehicle_id VARCHAR(20),
    latitude DOUBLE PRECISION,
    longitude DOUBLE PRECISION,
    timestamp BIGINT
);

-- Indexes
CREATE INDEX idx_id ON vehicle_locations USING btree (id);
CREATE INDEX idx_vehicle_id ON vehicle_locations USING btree (vehicle_id);
CREATE INDEX idx_timestamp ON vehicle_locations USING btree (timestamp);