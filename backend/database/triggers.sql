CREATE FUNCTION update_users()
RETURNS trigger AS $update_users$
BEGIN 
  UPDATE users
  SET owner_of = array_prepend(NEW.id, owner_of)
  WHERE id = NEW.owner;
  
  RETURN NEW;
END;
$update_users$ LANGUAGE plpgsql;

CREATE FUNCTION update_restaurant_capacity()
RETURNS trigger  AS $update_restaurant_capacity$
DECLARE 
  i INT;
  total_guests INT;
  max_capacity INT[];
  working_hours_v timerange[];
  days_open_v INT[];
BEGIN
  SELECT array_position(days_open, NEW.day_int[1]) INTO i FROM restaurants WHERE id=NEW.restaurant_id;

  SELECT capacity, working_hours, days_open INTO max_capacity, working_hours_v, days_open_v FROM restaurants WHERE id=NEW.restaurant_id;

  SELECT SUM(num_guests) INTO total_guests FROM reservations WHERE day=NEW.day;
  
  total_guests := total_guests;
  max_capacity := max_capacity;
  working_hours_v := working_hours_v;
  days_open_v := days_open_v;

  IF days_open_v[i] = NEW.day_int[1] THEN
    IF working_hours_v[i] @> NEW.h_from AND working_hours_v[i] @> NEW.h_to THEN
      IF total_guests > max_capacity[i] - NEW.num_guests THEN 
        RAISE EXCEPTION 'Not enough capacity';
      ELSE
        RETURN NEW;
      END IF;
    ELSE
      RAISE EXCEPTION 'Time out of working hours range';
    END IF;
  ELSE
    RAISE EXCEPTION 'Restaurant not open';
  END IF;
END;
$update_restaurant_capacity$ LANGUAGE plpgsql;

CREATE TRIGGER update_users_on_inserted_restaurant AFTER INSERT ON restaurants
FOR EACH ROW EXECUTE FUNCTION update_users();

CREATE TRIGGER update_restaurant_capacity_on_inserted_reservation BEFORE INSERT ON reservations
FOR EACH ROW EXECUTE FUNCTION update_restaurant_capacity();

CREATE TRIGGER update_restaurant_capacity_on_update_reservation BEFORE UPDATE ON reservations
FOR EACH ROW EXECUTE FUNCTION update_restaurant_capacity();
