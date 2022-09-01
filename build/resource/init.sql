create table if not exists items
(
    item_id bigint primary key,
    type varchar(255) not null,
    shortcut varchar(255) not null
);

create table if not exists storage_pairs
(
    storage_pair_id bigint primary key,
    value numeric not null,
    item_id bigint not null unique,
    foreign key (item_id) references items on delete cascade on update cascade
);

create table if not exists portfolios_save
(
    portfolio_save_id bigint primary key,
    name varchar(255) not null,
    storage_pair_id bigint,
    foreign key (storage_pair_id) references storage_pairs on delete cascade on update cascade
);

create table if not exists portfolios_low_risk
(
    portfolio_low_risk_id bigint primary key,
    name varchar(255) not null,
    storage_pair_id bigint,
    foreign key (storage_pair_id) references storage_pairs on delete cascade on update cascade
);

create table if not exists portfolios_high_risk
(
    portfolio_high_risk bigint primary key,
    name varchar(255) not null,
    storage_pair_id bigint,
    foreign key (storage_pair_id) references storage_pairs on delete cascade on update cascade
);

create table if not exists tabs
(
    tab_id bigint primary key,
    user_id varchar(255) not null,
    portfolio_save_id bigint unique,
    portfolio_low_risk_id bigint unique,
    portfolio_high_risk_id bigint unique,
    foreign key (portfolio_save_id) references portfolios_save on delete cascade on update cascade,
    foreign key (portfolio_low_risk_id) references portfolios_low_risk on delete cascade on update cascade,
    foreign key (portfolio_high_risk_id) references portfolios_high_risk on delete cascade on update cascade
)
