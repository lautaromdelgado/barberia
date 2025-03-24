-- Crear la base de datos
CREATE DATABASE IF NOT EXISTS barbershop_system;
USE barbershop_system;

-- Eliminar tablas si existen (para evitar errores en recreación)
DROP TABLE IF EXISTS haircuts;
DROP TABLE IF EXISTS barbershop_employees;
DROP TABLE IF EXISTS barbershops;
DROP TABLE IF EXISTS users;

-- Crear la tabla de usuarios (tanto dueños como empleados)
CREATE TABLE users (
  id INT AUTO_INCREMENT PRIMARY KEY,
  nombre VARCHAR(50) NOT NULL,
  apellido VARCHAR(50) NOT NULL,
  dni VARCHAR(255) NOT NULL UNIQUE,
  correo VARCHAR(100) UNIQUE, -- Ahora es opcional (puede ser NULL)
  rol ENUM('owner', 'employee') NOT NULL,
  verified TINYINT NOT NULL DEFAULT 0,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB;

-- Crear la tabla de barberías
CREATE TABLE barbershops (
  id INT AUTO_INCREMENT PRIMARY KEY,
  owner_id INT NOT NULL,
  nombre VARCHAR(100) NOT NULL,
  direccion VARCHAR(200) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT fk_barbershop_owner FOREIGN KEY (owner_id) REFERENCES users(id)
    ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB;

-- Crear la tabla que relaciona empleados con una barbería
CREATE TABLE barbershop_employees (
  id INT AUTO_INCREMENT PRIMARY KEY,
  barbershop_id INT NOT NULL,
  user_id INT NOT NULL,
  comision_porcentaje_default DECIMAL(5,2) NOT NULL DEFAULT 0, -- % de comisión por cada corte
  base_salary DECIMAL(10,2) NOT NULL DEFAULT 0, -- Salario base fijo del empleado
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT fk_bshopemp_barbershop FOREIGN KEY (barbershop_id) REFERENCES barbershops(id)
    ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT fk_bshopemp_user FOREIGN KEY (user_id) REFERENCES users(id)
    ON DELETE CASCADE ON UPDATE CASCADE,
  UNIQUE KEY uniq_barbershop_employee (barbershop_id, user_id)
) ENGINE=InnoDB;

-- Crear la tabla de cortes de pelo (servicios realizados)
CREATE TABLE haircuts (
  id INT AUTO_INCREMENT PRIMARY KEY,
  barbershop_id INT NOT NULL,
  user_id INT NOT NULL, -- Usuario que realizó el corte (puede ser dueño o empleado)
  realizado_en DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  monto_total DECIMAL(10,2) NOT NULL,
  comision_aplicada DECIMAL(10,2) NOT NULL DEFAULT 0, -- Se calcula solo para empleados
  porcentaje_comision DECIMAL(5,2) NOT NULL DEFAULT 0,  -- Se almacena para preservar el histórico
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT fk_haircuts_barbershop FOREIGN KEY (barbershop_id) REFERENCES barbershops(id)
    ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT fk_haircuts_user FOREIGN KEY (user_id) REFERENCES users(id)
    ON DELETE CASCADE ON UPDATE CASCADE,
  INDEX idx_realizado_en (realizado_en),
  INDEX idx_barbershop_realizado (barbershop_id, realizado_en)
) ENGINE=InnoDB;

-- Crear tabla para manejar los tokens de verificación de cuenta y recuperación de contraseña
CREATE TABLE user_tokens (
  id INT AUTO_INCREMENT PRIMARY KEY,
  user_id INT NOT NULL,
  token VARCHAR(255) NOT NULL UNIQUE,
  type ENUM('verification', 'password_reset') NOT NULL, -- Indica si es para verificación o recuperación
  expires_at DATETIME NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT fk_user_tokens_user FOREIGN KEY (user_id) REFERENCES users(id)
    ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB;
