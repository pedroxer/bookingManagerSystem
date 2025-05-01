-- First drop tables with foreign keys pointing to other tables
DROP TABLE IF EXISTS booking_service.parking_bookings;
DROP TABLE IF EXISTS booking_service.booking;
DROP TABLE IF EXISTS resource_service.items;
DROP TABLE IF EXISTS auth.role_permissions;

-- Then droIF EXISTS p tables that are only referenced by others
DROP TABLE IF EXISTS resource_service.workplace;
DROP TABLE IF EXISTS resource_service.item_conditions;
DROP TABLE IF EXISTS resource_service.parking_spaces;
DROP TABLE IF EXISTS auth.users;
DROP TABLE IF EXISTS auth.positions;
DROP TABLE IF EXISTS auth.permissions;
DROP TABLE IF EXISTS auth.user_roles;