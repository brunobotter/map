-- Ativar extensão PostGIS
CREATE EXTENSION IF NOT EXISTS postgis;

-- Tabela de eventos geográficos
CREATE TABLE events (
    id UUID PRIMARY KEY,
    type TEXT NOT NULL,
    severity INTEGER,
    source TEXT,

    -- localização geográfica
    location GEOGRAPHY(POINT, 4326) NOT NULL,

    created_at TIMESTAMP DEFAULT NOW()
);

-- Índice geográfico para buscas rápidas
CREATE INDEX idx_events_location
ON events
USING GIST(location);