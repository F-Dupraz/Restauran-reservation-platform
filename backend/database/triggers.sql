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
  remaining_cap INT;
  working_hours_v timerange[];
  days_open_v INT[];
BEGIN
  SELECT array_position(days_open, NEW.day[1]) INTO i FROM restaurants WHERE id=NEW.restaurant_id;

  SELECT capacity[i], working_hours, days_open INTO remaining_cap, working_hours_v, days_open_v FROM restaurants WHERE id=NEW.restaurant_id;

  remaining_cap := remaining_cap - NEW.num_guests;
  working_hours_v := working_hours_v;
  days_open_v := days_open_v;

  IF days_open_v[i] = NEW.day[1] THEN
    IF working_hours_v[i] @> NEW.h_from AND working_hours_v[i] @> NEW.h_to THEN
      IF remaining_cap >= 0 THEN 
        UPDATE restaurants
        SET capacity[i] = capacity[i] - NEW.num_guests
        WHERE id=NEW.restaurant_id;

        RETURN NEW;
      ELSE
        DELETE FROM reservations WHERE id=NEW.id;

        RETURN "There is not enough capacity to make the reservation";
      END IF;
    ELSE
      DELETE FROM reservations WHERE id=NEW.id;

      RETURN "Error 2";
    END IF;
  ELSE
    DELETE FROM reservations WHERE id=NEW.id;

    RETURN "Error 1";
  END IF;
END;
$update_restaurant_capacity$ LANGUAGE plpgsql;

CREATE TRIGGER update_users_on_inserted_restaurant AFTER INSERT ON restaurants
FOR EACH ROW EXECUTE FUNCTION update_users();

CREATE TRIGGER update_restaurant_capacity_on_inserted_reservation AFTER INSERT ON reservations
FOR EACH ROW EXECUTE FUNCTION update_restaurant_capacity();

CREATE TRIGGER update_restaurant_capacity_on_update_reservation AFTER UPDATE ON reservations
FOR EACH ROW EXECUTE FUNCTION update_restaurant_capacity();
