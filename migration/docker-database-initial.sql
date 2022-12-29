create table receitas(
    id serial primary key,
    nome varchar,
    ingredientes varchar
);

INSERT INTO receitas(nome, ingredientes) VALUES
('Burguer Popeye', 'Hamburger de ervilha com espinafre, aveia, farinha de trigo e de rosca.'),
('Burguer Hulk', 'Hamburguer de couve refogada com farinha de aveia, de trigo e de rosca, tempero a gosto.');