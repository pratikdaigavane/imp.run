-- Create a keyspace
CREATE KEYSPACE IF NOT EXISTS url_shortener WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : '1' };

-- Create a table
CREATE TABLE IF NOT EXISTS url_shortener.urls
(
    short_code text PRIMARY KEY,
    url        text,
    created_at timestamp
);