pragma foreign_keys = on;

create table beers (
	ID        integer not null primary key,
	Name      varchar not null,
	Brewery   varchar,
	ABV       integer,
	ShortDesc varchar,
	CreatedAt integer not null default current_timestamp,
	UpdatedAt integer not null default current_timestamp,
	DeletedAt integer default null
);

-- Set updatedAt when any field is touched
create trigger beers_UpdatedAt after update on beers
        begin
                update beers set UpdatedAt = current_timestamp where title=old.title;
        end;

-- Seed some data
insert into beers (Name) values
	("La Trappe Dubbel"),
	("Gouden Carolus Classic")
;
