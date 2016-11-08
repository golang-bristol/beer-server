pragma foreign_keys = on;

create table beers (
	ID        integer not null primary key,
	Name      varchar not null,
	Brewery   integer,
	ABV       integer,
	ShortDesc varchar,
	CreatedAt integer not null default current_timestamp,
	UpdatedAt integer not null default current_timestamp,
	DeletedAt integer default null
);

create table breweries (
	ID        integer not null primary key,
	Name      varchar not null,
	Location  varchar not null,
	CreatedAt integer not null default current_timestamp,
	UpdatedAt integer not null default current_timestamp,
	DeletedAt integer default null
);

create table reviews (
	ID        integer not null primary key,
	Beer      integer not null,
	Name      varchar not null,
	Score     integer not null,
	Text      varchar,
	CreatedAt integer not null default current_timestamp,
	UpdatedAt integer not null default current_timestamp,
	DeletedAt integer default null
);

-- Set updatedAt when any field is touched
create trigger beers_UpdatedAt after update on beers
        begin
                update beers set UpdatedAt = current_timestamp where title=old.title;
        end;
create trigger reviews_UpdatedAt after update on reviews
        begin
                update reviews set UpdatedAt = current_timestamp where title=old.title;
        end;

-- Seed some data
insert into beers (Name) values
	("La Trappe Dubbel"),
	("Gouden Carolus Classic")
;
