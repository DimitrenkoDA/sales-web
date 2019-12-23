create table subs(sub_id int, sub_name varchar(30), address varchar(30), phone varchar(11), note varchar(100));
alter table subs add constraint sub_id_PK primary key(sub_id);

create table statuses(status_id int, status_group int, status_name varchar(30), note varchar(100));
alter table statuses add constraint status_id_PK primary key(status_id);
alter table statuses add constraint status_id_FK foreign key(status_group) references statuses(status_id);

create table product_groups(pg_id int, pg_name varchar(30), note varchar(100));
alter table product_groups add constraint pg_id_PK primary key(pg_id);

create table deals(n_deal int, data_start date, data_finish date, prim varchar(100) default 'no comments', sub_id int);
alter table deals add constraint n_deal_PK primary key(n_deal);
alter table deals add constraint sub_id_FK foreign key(sub_id) references subs(sub_id);

create table dealers(n_dealer int, dealer_name varchar(30), address varchar(30), phone varchar(11), status_id int, note varchar(100));
alter table dealers add constraint n_dealer_PK primary key(n_dealer);
alter table dealers add constraint status_id_FK2 foreign key(status_id) references statuses(status_id);

create table products(prod_id int, prod_name varchar(30), pincode int, data_iz date, pg_id int, status_id int, note varchar(100));
alter table products add constraint prod_id_PK primary key(prod_id);
alter table products add constraint pg_id_FK foreign key(PG_id) references product_groups(pg_id);
alter table products add constraint status_id_FK3 foreign key(status_id) references statuses(status_id);

create table pricelist(prod_id int, dat date, price int, note varchar(100));
alter table pricelist add constraint prod_id_dat_PK primary key(prod_id,dat);
alter table pricelist add constraint prod_id_FK foreign key(prod_id) references products(prod_id);

create table salemans(man_code int, saleman_name varchar(30), card_code int, n_dealer int, status_id int, note varchar(100), condition varchar (6));
alter table salemans add constraint man_code_PK primary key(man_code);
alter table salemans add constraint n_dealer_FK foreign key(n_dealer) references dealers(N_dealer);
alter table salemans add constraint status_id_FK4 foreign key(status_id) references statuses(status_id);

create table salemaps(map_id int, prod_id int, dat date, sub_id int, man_code int, quantity int, sale_dat date, note varchar(100));
alter table salemaps add constraint map_id_PK primary key(map_id);
alter table salemaps add constraint prod_id_dat_FK foreign key(prod_id,dat) references pricelist(prod_id,dat);
alter table salemaps add constraint sub_id_FK2 foreign key(sub_id) references subs(sub_id);
alter table salemaps add constraint man_code_FK foreign key(man_code) references salemans(man_code);

insert into subs(sub_id, sub_name, address, phone, note) values (1000, 'Boris Razor', '10 Lenin st 43 Russia', 88006567385, null);
insert into subs(sub_id, sub_name, address, phone, note) values (1001, 'Anton Ruzke', '12 Marks st 31 Russia', 89035789896, null);
insert into subs(sub_id, sub_name, address, phone, note) values (1002, 'Ivan Gamaz', '41 Mira st 12 Russia', 89438550306, null);
insert into subs(sub_id, sub_name, address, phone, note) values (1003, 'Ilya Maddyson', '5 Svyazi st 24 Russia', 89007568354, null);

insert into product_groups(pg_id, pg_name, note) values (111, 'Electronics', null);
insert into product_groups(pg_id, pg_name, note) values (112, 'Appliances', null);
insert into product_groups(pg_id, pg_name, note) values (113, 'Lighting', null);
insert into product_groups(pg_id, pg_name, note) values (114, 'Computers', null);
insert into product_groups(pg_id, pg_name, note) values (115, 'Gadgets', null);

insert into statuses(status_id, status_group, status_name, note) values(1, null,'ProductsGroup','group');
insert into statuses(status_id, status_group, status_name, note) values(2, null,'DealersGroup','group');
insert into statuses(status_id, status_group, status_name, note) values(3, null,'SalemansGroup','group');

insert into statuses(status_id, status_group, status_name, note) values(4, 1,'New', null);
insert into statuses(status_id, status_group, status_name, note) values(5, 1,'Secondhand', null);

insert into statuses(status_id, status_group, status_name, note) values(6, 2,'Available', null);
insert into statuses(status_id, status_group, status_name, note) values(7, 2,'Unavailable', null);

insert into statuses(status_id, status_group, status_name, note) values(8, 3,'Free', null);
insert into statuses(status_id, status_group, status_name, note) values(9, 3,'Busy', null);

insert into products(prod_id, prod_name, pincode, data_iz, pg_id, status_id, note) values(1, 'Laptop Acer A517-51G-32F5', 830543, '03.15.2017', 114, 4, null);
insert into products(prod_id, prod_name, pincode, data_iz, pg_id, status_id, note) values(2, 'Laptop Asus FX753VD-GC448T', 370192, '01.20.2018', 114, 4, null);
insert into products(prod_id, prod_name, pincode, data_iz, pg_id, status_id, note) values(3, 'TV LG 49LK6200', 280694, '02.11.2010', 111, 5, null);
insert into products(prod_id, prod_name, pincode, data_iz, pg_id, status_id, note) values(4, 'Fridge Haier A2F637CGG', 711049, '12.18.2013', 112, 4, null);
insert into products(prod_id, prod_name, pincode, data_iz, pg_id, status_id, note) values(5, 'Washer Ariston WMTG 722', 201575, '11.28.2003', 112, 5, null);
insert into products(prod_id, prod_name, pincode, data_iz, pg_id, status_id, note) values(6, 'Stove Gorenje K634WA', 364205, '07.30.2007', 112, 5, null);
insert into products(prod_id, prod_name, pincode, data_iz, pg_id, status_id, note) values(7, 'Phone Samsung Galaxy S10', 774951, '06.20.2018', 115, 5, null);
insert into products(prod_id, prod_name, pincode, data_iz, pg_id, status_id, note) values(8, 'Phone Honor 10', 256719, '08.22.2018', 115, 4, null);
insert into products(prod_id, prod_name, pincode, data_iz, pg_id, status_id, note) values(9, 'Lamp LED Endever', 556342, '06.14.2017', 113, 4, null);
insert into products(prod_id, prod_name, pincode, data_iz, pg_id, status_id, note) values(10, 'Audio Mini LG DM5660K', 972182, '10.14.2016', 111, 5, null);
insert into products(prod_id, prod_name, pincode, data_iz, pg_id, status_id, note) values(11, 'PC HP 460-p231ur', 763553, '05.19.2017', 114, 4, null);
insert into products(prod_id, prod_name, pincode, data_iz, pg_id, status_id, note) values(12, 'Combo Sven Challenge', 838039, '09.09.2015', 115, 4, null);
insert into products(prod_id, prod_name, pincode, data_iz, pg_id, status_id, note) values(13, 'Rudder Thrustmaster T150', 966787, '01.01.2010', 115, 5, null);
insert into products(prod_id, prod_name, pincode, data_iz, pg_id, status_id, note) values(14, 'Air HEC 09HTC03', 465087, '02.24.2013', 111, 5, null);
insert into products(prod_id, prod_name, pincode, data_iz, pg_id, status_id, note) values(15, 'Moulinex CE501132', 268639, '07.04.2016', 112, 4, null);
insert into products(prod_id, prod_name, pincode, data_iz, pg_id, status_id, note) values(16, 'Camera Fujifilm Instax', 461991, '10.09.2017', 115, 5, null);
insert into products(prod_id, prod_name, pincode, data_iz, pg_id, status_id, note) values(17, 'Blu-Ray Philips BDP3290B', 227350, '12.19.2009', 115, 5, null);
insert into products(prod_id, prod_name, pincode, data_iz, pg_id, status_id, note) values(18, 'Recorder Prology MPN-450', 616654, '03.29.2019', 111, 4, null);
insert into products(prod_id, prod_name, pincode, data_iz, pg_id, status_id, note) values(19, 'Chandelier REGENBOGEN', 492552, '04.11.2015', 113, 4, null);
insert into products(prod_id, prod_name, pincode, data_iz, pg_id, status_id, note) values(20, 'Lamp EGLO TOWNSHEND 32918', 504812, '07.01.2017', 113, 5, null);
insert into products(prod_id, prod_name, pincode, data_iz, pg_id, status_id, note) values(21, 'Fan Tefal VF6210F0', 732147, '09.26.2015', 112, 4, null);
insert into products(prod_id, prod_name, pincode, data_iz, pg_id, status_id, note) values(22, 'Printer Canon PIXMA', 711013, '11.13.2017', 111, 5, null);
insert into products(prod_id, prod_name, pincode, data_iz, pg_id, status_id, note) values(23, 'Drone DJI Mavic 2', 956643, '07.16.2015', 115, 4, null);
insert into products(prod_id, prod_name, pincode, data_iz, pg_id, status_id, note) values(24, 'Ventilation Kuppersberg F', 621124, '08.08.2011', 111, 5, null);

insert into dealers(n_dealer, dealer_name, address, phone, status_id, note) values (1, 'Alladin', '1 Baldin st 16 Russia', 89268567274, 7, null);
insert into dealers(n_dealer, dealer_name, address, phone, status_id, note) values (2, 'Don Carleone', '4 Sicily st 99 Italy', 89006364833, 7, null);
insert into dealers(n_dealer, dealer_name, address, phone, status_id, note) values (3, 'Goblin', '12/2 2ch st 13 Russia', 89994562869, 6, null);

insert into deals(n_deal, data_start, data_finish, prim, sub_id) values(1, '01.08.1999', null, 'no comments', 1000);
insert into deals(n_deal, data_start, data_finish, prim, sub_id) values(2, '04.11.2003', '06.21.2006', 'no comments', 1000);
insert into deals(n_deal, data_start, data_finish, prim, sub_id) values(3, '01.08.2017', null, 'no comments', 1000);
insert into deals(n_deal, data_start, data_finish, prim, sub_id) values(4, '07.08.2010', '08.08.2010', 'no comments', 1001);
insert into deals(n_deal, data_start, data_finish, prim, sub_id) values(5, '01.05.2011', '03.28.2014', 'no comments', 1001);
insert into deals(n_deal, data_start, data_finish, prim, sub_id) values(6, '09.12.1995', null, 'no comments', 1002);
insert into deals(n_deal, data_start, data_finish, prim, sub_id) values(7, '12.12.1992', '01.11.1995', 'no comments', 1003);
insert into deals(n_deal, data_start, data_finish, prim, sub_id) values(8, '12.12.1998', '01.11.2009', 'no comments', 1003);
insert into deals(n_deal, data_start, data_finish, prim, sub_id) values(9, '7.22.2001', '04.30.2004', 'no comments', 1003);
insert into deals(n_deal, data_start, data_finish, prim, sub_id) values(10, '5.13.2012', null, 'no comments', 1003);

insert into pricelist(prod_id, dat, price, note) values(24,'04.04.2017', 1399, null);
insert into pricelist(prod_id, dat, price, note) values(23,'03.01.2019', 9919, null);
insert into pricelist(prod_id, dat, price, note) values(22,'04.29.2018', 7990, null);
insert into pricelist(prod_id, dat, price, note) values(21,'01.09.2016', 3990, null);
insert into pricelist(prod_id, dat, price, note) values(20,'10.22.2017', 3090, null);
insert into pricelist(prod_id, dat, price, note) values(19,'02.28.2018', 3949, null);
insert into pricelist(prod_id, dat, price, note) values(18,'09.06.2017', 9990, null);
insert into pricelist(prod_id, dat, price, note) values(17,'02.19.2017', 7990, null);
insert into pricelist(prod_id, dat, price, note) values(16,'06.06.2016', 8490, null);
insert into pricelist(prod_id, dat, price, note) values(15,'04.21.2018', 6490, null);
insert into pricelist(prod_id, dat, price, note) values(14,'08.04.2016', 1499, null);
insert into pricelist(prod_id, dat, price, note) values(13,'01.17.2017', 2199, null);
insert into pricelist(prod_id, dat, price, note) values(12,'07.17.2017', 1590, null);
insert into pricelist(prod_id, dat, price, note) values(11,'07.20.2017', 2999, null);
insert into pricelist(prod_id, dat, price, note) values(10,'10.02.2016', 1299, null);
insert into pricelist(prod_id, dat, price, note) values(9,'04.08.2017', 490, null);
insert into pricelist(prod_id, dat, price, note) values(8,'01.30.2017', 2499, null);
insert into pricelist(prod_id, dat, price, note) values(7,'06.08.2017', 6899, null);
insert into pricelist(prod_id, dat, price, note) values(6,'06.12.2017', 3299, null);
insert into pricelist(prod_id, dat, price, note) values(5,'05.05.2018', 3199, null);
insert into pricelist(prod_id, dat, price, note) values(4,'12.13.2017', 5999, null);
insert into pricelist(prod_id, dat, price, note) values(3,'07.10.2018', 3599, null);
insert into pricelist(prod_id, dat, price, note) values(2,'07.10.2018', 7299, null);
insert into pricelist(prod_id, dat, price, note) values(1,'07.10.2018', 4999, null);
insert into pricelist(prod_id, dat, price, note) values(19,'10.25.2017', 3749, null);
insert into pricelist(prod_id, dat, price, note) values(9,'03.16.2018', 1490, null);
insert into pricelist(prod_id, dat, price, note) values(15,'01.23.2016', 6090, null);
insert into pricelist(prod_id, dat, price, note) values(22,'01.14.2016', 8990, null);
insert into pricelist(prod_id, dat, price, note) values(24,'05.16.2016', 1359, null);
insert into pricelist(prod_id, dat, price, note) values(12,'10.16.2016', 1790, null);
insert into pricelist(prod_id, dat, price, note) values(11,'09.19.2017', 3100, null);
insert into pricelist(prod_id, dat, price, note) values(5,'02.01.2018', 3299, null);
insert into pricelist(prod_id, dat, price, note) values(4,'01.11.2017', 5599, null);
insert into pricelist(prod_id, dat, price, note) values(14,'02.25.2016', 1599, null);
insert into pricelist(prod_id, dat, price, note) values(11,'10.07.2017', 2499, null);
insert into pricelist(prod_id, dat, price, note) values(10,'03.14.2019', 1199, null);

insert into salemans(man_code, saleman_name, card_code, n_dealer, status_id, note, condition) values(11, 'Andrew', 3219, 1, 8, null, 'common');
insert into salemans(man_code, saleman_name, card_code, n_dealer, status_id, note, condition) values(12, 'Carl', 7371, 3, 9, null, 'common');
insert into salemans(man_code, saleman_name, card_code, n_dealer, status_id, note, condition) values(13, 'Ivan', 6371, 2, 8, null, 'common');
insert into salemans(man_code, saleman_name, card_code, n_dealer, status_id, note, condition) values(14, 'Leo', 8827, 1, 9, null, 'common');
insert into salemans(man_code, saleman_name, card_code, n_dealer, status_id, note, condition) values(15, 'Simon', 4282, 3, 8, null, 'common');
insert into salemans(man_code, saleman_name, card_code, n_dealer, status_id, note, condition) values(16, 'Jack', 2184, 1, 9, null, 'VIP');
insert into salemans(man_code, saleman_name, card_code, n_dealer, status_id, note, condition) values(17, 'Sandro', 0912, 2, 8, null, 'VIP');
insert into salemans(man_code, saleman_name, card_code, n_dealer, status_id, note, condition) values(18, 'Paul', 1243, 3, 9, null, 'VIP');
insert into salemans(man_code, saleman_name, card_code, n_dealer, status_id, note, condition) values(19, 'Linda', 1241, 2, 8, null, 'VIP');
insert into salemans(man_code, saleman_name, card_code, n_dealer, status_id, note, condition) values(20, 'Frank', 8242, 1, 9, null, 'VIP');

insert into salemaps(map_id, prod_id, dat, sub_id, man_code, quantity, sale_dat, note) values(100, 24, '04.04.2017', 1000, 11, 3, '07.24.2016', null);
insert into salemaps(map_id, prod_id, dat, sub_id, man_code, quantity, sale_dat, note) values(101, 23, '03.01.2019', 1002, 13, 1, '05.05.2017', null);
insert into salemaps(map_id, prod_id, dat, sub_id, man_code, quantity, sale_dat, note) values(102, 22, '04.29.2018', 1001, 15, 5, '11.09.2017', null);
insert into salemaps(map_id, prod_id, dat, sub_id, man_code, quantity, sale_dat, note) values(103, 21, '01.09.2016', 1003, 12, 8, '02.16.2019', null);
insert into salemaps(map_id, prod_id, dat, sub_id, man_code, quantity, sale_dat, note) values(104, 20, '10.22.2017', 1000, 14, 1, '10.24.2016', null);
insert into salemaps(map_id, prod_id, dat, sub_id, man_code, quantity, sale_dat, note) values(105, 19, '02.28.2018', 1000, 13, 12, '03.25.2016', null);
insert into salemaps(map_id, prod_id, dat, sub_id, man_code, quantity, sale_dat, note) values(106, 18, '09.06.2017', 1001, 14, 4, '02.28.2017', null);
insert into salemaps(map_id, prod_id, dat, sub_id, man_code, quantity, sale_dat, note) values(107, 17, '02.19.2017', 1002, 12, 3, '12.26.2017', null);
insert into salemaps(map_id, prod_id, dat, sub_id, man_code, quantity, sale_dat, note) values(108, 16, '06.06.2016', 1003, 11, 2, '12.21.2016', null);
insert into salemaps(map_id, prod_id, dat, sub_id, man_code, quantity, sale_dat, note) values(109, 15, '04.21.2018', 1000, 15, 12, '10.03.2017', null);
insert into salemaps(map_id, prod_id, dat, sub_id, man_code, quantity, sale_dat, note) values(110, 14, '08.04.2016', 1002, 13, 3, '12.17.2018', null);
insert into salemaps(map_id, prod_id, dat, sub_id, man_code, quantity, sale_dat, note) values(111, 13, '01.17.2017', 1001, 11, 12, '12.09.2016', null);
insert into salemaps(map_id, prod_id, dat, sub_id, man_code, quantity, sale_dat, note) values(112, 12, '07.17.2017', 1003, 14, 1, '04.08.2018', null);
insert into salemaps(map_id, prod_id, dat, sub_id, man_code, quantity, sale_dat, note) values(113, 11, '07.20.2017', 1002, 15, 12, '01.01.2019', null);
insert into salemaps(map_id, prod_id, dat, sub_id, man_code, quantity, sale_dat, note) values(114, 10, '10.02.2016', 1000, 12, 2, '01.13.2018', null);
insert into salemaps(map_id, prod_id, dat, sub_id, man_code, quantity, sale_dat, note) values(115, 9, '04.08.2017', 1003, 13, 21, '02.25.2017', null);
insert into salemaps(map_id, prod_id, dat, sub_id, man_code, quantity, sale_dat, note) values(116, 8, '01.30.2017', 1001, 11, 11, '03.03.2017', null);
insert into salemaps(map_id, prod_id, dat, sub_id, man_code, quantity, sale_dat, note) values(117, 7, '06.08.2017', 1002, 12, 7, '01.28.2018', null);
insert into salemaps(map_id, prod_id, dat, sub_id, man_code, quantity, sale_dat, note) values(118, 6, '06.12.2017', 1003, 13, 3, '05.28.2018', null);
insert into salemaps(map_id, prod_id, dat, sub_id, man_code, quantity, sale_dat, note) values(119, 5, '05.05.2018', 1000, 14, 12, '06.28.2018', null);
insert into salemaps(map_id, prod_id, dat, sub_id, man_code, quantity, sale_dat, note) values(120, 4, '12.13.2017', 1000, 15, 9, '05.19.2017', null);
insert into salemaps(map_id, prod_id, dat, sub_id, man_code, quantity, sale_dat, note) values(121, 3, '07.10.2018', 1001, 11, 2, '01.04.2019', null);
insert into salemaps(map_id, prod_id, dat, sub_id, man_code, quantity, sale_dat, note) values(122, 2, '07.10.2018', 1002, 12, 21, '01.16.2016', null);
insert into salemaps(map_id, prod_id, dat, sub_id, man_code, quantity, sale_dat, note) values(123, 1, '07.10.2018', 1003, 13, 4, '10.06.2017', null);
insert into salemaps(map_id, prod_id, dat, sub_id, man_code, quantity, sale_dat, note) values(124, 19, '10.25.2017', 1001, 14, 12, '08.28.2018', null);
insert into salemaps(map_id, prod_id, dat, sub_id, man_code, quantity, sale_dat, note) values(125, 9, '03.16.2018', 1002, 15, 2, '06.07.2018', null);
insert into salemaps(map_id, prod_id, dat, sub_id, man_code, quantity, sale_dat, note) values(126, 15, '01.23.2016', 1003, 11, 12, '01.29.2019', null);
insert into salemaps(map_id, prod_id, dat, sub_id, man_code, quantity, sale_dat, note) values(127, 22, '01.14.2016', 1000, 12, 1, '03.06.2016', null);
insert into salemaps(map_id, prod_id, dat, sub_id, man_code, quantity, sale_dat, note) values(128, 24, '05.16.2016', 1001, 13, 6, '06.18.2018', null);
insert into salemaps(map_id, prod_id, dat, sub_id, man_code, quantity, sale_dat, note) values(129, 12, '10.16.2016', 1002, 14, 1, '06.27.2017', null);
insert into salemaps(map_id, prod_id, dat, sub_id, man_code, quantity, sale_dat, note) values(130, 11, '09.19.2017', 1003, 15, 8, '06.19.2018', null);
insert into salemaps(map_id, prod_id, dat, sub_id, man_code, quantity, sale_dat, note) values(131, 5, '02.01.2018', 1000, 11, 18, '02.20.2018', null);
insert into salemaps(map_id, prod_id, dat, sub_id, man_code, quantity, sale_dat, note) values(132, 4, '01.11.2017', 1001, 12, 1, '07.25.2017', null);
insert into salemaps(map_id, prod_id, dat, sub_id, man_code, quantity, sale_dat, note) values(133, 14, '02.25.2016', 1002, 13, 16, '08.22.2016', null);
insert into salemaps(map_id, prod_id, dat, sub_id, man_code, quantity, sale_dat, note) values(134, 11, '10.07.2017', 1003, 14, 5, '04.21.2018', null);
insert into salemaps(map_id, prod_id, dat, sub_id, man_code, quantity, sale_dat, note) values(135, 10,'03.14.2019', 1000, 15, 2, '12.17.2018', null);
insert into salemaps(map_id, prod_id, dat, sub_id, man_code, quantity, sale_dat, note) values(136, 10,'10.02.2016', 1001, 11, 5, '11.01.2017', null);
insert into salemaps(map_id, prod_id, dat, sub_id, man_code, quantity, sale_dat, note) values(137, 9, '04.08.2017', 1002, 12, 9, '07.26.2017', null);
insert into salemaps(map_id, prod_id, dat, sub_id, man_code, quantity, sale_dat, note) values(138, 18, '09.06.2017', 1003, 13, 5, '04.09.2018', null);
insert into salemaps(map_id, prod_id, dat, sub_id, man_code, quantity, sale_dat, note) values(139, 13, '01.17.2017', 1000, 14, 2, '10.01.2019', null);
insert into salemaps(map_id, prod_id, dat, sub_id, man_code, quantity, sale_dat, note) values(140, 8, '01.30.2017', 1001, 15, 7, '04.27.2018', null);
insert into salemaps(map_id, prod_id, dat, sub_id, man_code, quantity, sale_dat, note) values(141, 24, '04.04.2017', 1002, 11, 1, '01.16.2018', null);
insert into salemaps(map_id, prod_id, dat, sub_id, man_code, quantity, sale_dat, note) values(142, 21, '01.09.2016', 1003, 12, 8, '05.03.2018', null);
insert into salemaps(map_id, prod_id, dat, sub_id, man_code, quantity, sale_dat, note) values(143, 7, '06.08.2017', 1000, 13, 9, '09.08.2017', null);
insert into salemaps(map_id, prod_id, dat, sub_id, man_code, quantity, sale_dat, note) values(144, 3, '07.10.2018', 1001, 14, 1, '03.15.2019', null);
insert into salemaps(map_id, prod_id, dat, sub_id, man_code, quantity, sale_dat, note) values(145, 12, '07.17.2017', 1002, 15, 3, '01.30.2017', null);
insert into salemaps(map_id, prod_id, dat, sub_id, man_code, quantity, sale_dat, note) values(146, 6, '06.12.2017', 1003, 11, 6, '10.15.2016', null);
insert into salemaps(map_id, prod_id, dat, sub_id, man_code, quantity, sale_dat, note) values(147, 5, '05.05.2018', 1000, 12, 1, '10.27.2017', null);
insert into salemaps(map_id, prod_id, dat, sub_id, man_code, quantity, sale_dat, note) values(148, 3,'07.10.2018', 1001, 13, 4, '07.03.2018', null);
insert into salemaps(map_id, prod_id, dat, sub_id, man_code, quantity, sale_dat, note) values(149, 18, '09.06.2017', 1002, 14, 2, '03.31.2017', null);