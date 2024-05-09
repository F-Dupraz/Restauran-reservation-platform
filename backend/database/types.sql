-- COPIADO DE LA DOCUMENTACION OFICIAL DE POSTGRES --

CREATE FUNCTION time_subtype_diff(x TIME, y TIME) RETURNS FLOAT AS
'SELECT EXTRACT(EPOCH FROM (x - y))' LANGUAGE sql STRICT IMMUTABLE;

CREATE TYPE timerange AS RANGE (
  subtype = TIME,
  subtype_diff = time_subtype_diff
);
