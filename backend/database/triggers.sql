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
BEGIN
  SELECT array_position(days_open, NEW.day) INTO i FROM restaurants WHERE id=NEW.restaurant_id;

  SELECT capacity[i] INTO remaining_cap FROM restaurants WHERE id=NEW.restaurant_id;

  remaining_cap := remaining_cap - NEW.num_guests;

  IF remaining_cap >= 0 THEN 
    UPDATE restaurants
    SET capacity[i] = capacity[i] - NEW.num_guests
    WHERE id=NEW.restaurant_id;

    RETURN NEW;
  ELSE
    RETURN "There is not enough capacity to make the reservation";
  END IF;

END;
$update_restaurant_capacity$ LANGUAGE plpgsql;

CREATE TRIGGER update_users_on_inserted_restaurant AFTER INSERT ON restaurants
FOR EACH ROW EXECUTE FUNCTION update_users();

CREATE TRIGGER update_restaurant_capacity_on_inserted_reservation AFTER INSERT ON reservations
FOR EACH ROW EXECUTE FUNCTION update_restaurant_capacity();
