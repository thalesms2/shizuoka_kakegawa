CREATE TABLE users (
  id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
  tenant_id int NOT NULL,
  deleted boolean DEFAULT false,
  created_at timestamptz DEFAULT now(),
  updated_at timestamptz DEFAULT now(),
  deleted_at timestamptz DEFAULT NULL,

  name varchar(255) NOT NULL,
  email varchar(255) UNIQUE NOT NULL,
  password varchar(255) NOT NULL,
  role varchar(255) NOT NULL,

  CONSTRAINT fk_tenant FOREIGN KEY (tenant_id) REFERENCES tenants(id),
);

CREATE TABLE transactions (
  id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
  tenant_id int NOT NULL,
  deleted boolean DEFAULT false,
  created_at timestamptz DEFAULT now(),
  updated_at timestamptz DEFAULT now(),
  deleted_at timestamptz DEFAULT NULL,

  description varchar(255),
  value int NOT NULL,
  operation varchar(255) NOT NULL,
  client_id int NOT NULL,
  invoice_id int REFERENCES invoices(id),
  service_id int REFERENCES services(id),

  CONSTRAINT fk_tenant FOREIGN KEY (tenant_id) REFERENCES tenants(id),
  CONSTRAINT fk_client FOREIGN KEY (client_id) REFERENCES clients(id),
);

CREATE TABLE invoices (
  id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
  tenant_id int NOT NULL,
  deleted boolean DEFAULT false,
  created_at timestamptz DEFAULT now(),
  updated_at timestamptz DEFAULT now(),
  deleted_at timestamptz DEFAULT NULL,

  description varchar(255),
  value int NOT NULL,
  balance int NOT NULL,
  due timestamptz DEFAULT NULL,
  canceled boolean DEFAULT false,
  canceled_at timestamptz DEFAULT NULL,
  client_id int,

  CONSTRAINT fk_tenant FOREIGN KEY (tenant_id) REFERENCES tenants(id),
  CONSTRAINT fk_client FOREIGN KEY (client_id) REFERENCES clients(id),
);

CREATE TABLE tenants (
  id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
  deleted boolean DEFAULT false,
  created_at timestamptz DEFAULT now(),
  updated_at timestamptz DEFAULT now(),
  deleted_at timestamptz DEFAULT NULL,

  name varchar(255),
);

CREATE TABLE clients (
  id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
  tenant_id int NOT NULL,
  deleted boolean DEFAULT false,
  created_at timestamptz DEFAULT now(),
  updated_at timestamptz DEFAULT now(),
  deleted_at timestamptz DEFAULT NULL,

  name varchar(255) NOT NULL,
  phone varchar(255) NOT NULL,

  CONSTRAINT fk_tenant FOREIGN KEY (tenant_id) REFERENCES tenants(id),
);

CREATE TABLE access_histories (
  id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
  tenant_id int NOT NULL,
  created_at timestamptz DEFAULT now(),

  user_id int NOT NULL,

  CONSTRAINT fk_tenant FOREIGN KEY (tenant_id) REFERENCES tenants(id),
  CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id),
);

CREATE TABLE services (
  id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
  tenant_id int NOT NULL,
  deleted boolean DEFAULT false,
  created_at timestamptz DEFAULT now(),
  updated_at timestamptz DEFAULT now(),
  deleted_at timestamptz DEFAULT NULL,

  description varchar(255),
  value int,
  balance int,
  canceled boolean DEFAULT false,
  canceled_at timestamptz DEFAULT NULL,
  client_id int,
  user_id int REFERENCES users(id),


  CONSTRAINT fk_tenant FOREIGN KEY (tenant_id) REFERENCES tenants(id),
  CONSTRAINT fk_client FOREIGN KEY (client_id) REFERENCES clients(id),
);
