-- Таблица для хранения складов
CREATE TABLE warehouses (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(), -- Уникальный идентификатор склада
    address TEXT NOT NULL -- Адрес склада
);

-- Таблица для хранения продуктов
CREATE TABLE products (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(), -- Уникальный идентификатор продукта
    name TEXT NOT NULL, -- Название продукта
    description TEXT, -- Описание продукта
    characteristics JSONB, -- Характеристики продукта в формате JSON
    weight FLOAT NOT NULL, -- Вес продукта
    barcode TEXT NOT NULL UNIQUE -- Штрих-код продукта (должен быть уникальным)
);

-- Таблица для хранения инвентаря
CREATE TABLE inventory (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(), -- Уникальный идентификатор записи инвентаря
    product_id UUID REFERENCES products(id) ON DELETE CASCADE, -- Идентификатор продукта
    warehouse_id UUID REFERENCES warehouses(id) ON DELETE CASCADE, -- Идентификатор склада
    quantity INT NOT NULL, -- Количество на складе
    price FLOAT NOT NULL, -- Цена продукта
    discount FLOAT -- Скидка на продукт
);

-- Таблица для хранения аналитики
CREATE TABLE analytics (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(), -- Уникальный идентификатор записи аналитики
    warehouse_id UUID REFERENCES warehouses(id) ON DELETE CASCADE, -- Идентификатор склада
    product_id UUID REFERENCES products(id) ON DELETE CASCADE, -- Идентификатор продукта
    sold_quantity INT NOT NULL, -- Количество проданных единиц
    total_amount FLOAT NOT NULL -- Общая сумма продаж
);

-- Индексы для повышения производительности
CREATE INDEX idx_inventory_product ON inventory(product_id);
CREATE INDEX idx_inventory_warehouse ON inventory(warehouse_id);
CREATE INDEX idx_analytics_warehouse ON analytics(warehouse_id);
CREATE INDEX idx_analytics_product ON analytics(product_id);