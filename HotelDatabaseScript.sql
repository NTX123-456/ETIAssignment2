CREATE database ETIASSG2_db
use ETIASSG2_db


CREATE TABLE Hotels (ID int AUTO_INCREMENT NOT NULL PRIMARY KEY,HotelName VARCHAR(30), HotelInfo VARCHAR(100), HotelAddr VARCHAR(100),HotelStar int, HotelAmenities VARCHAR(70),price int ,Country VARCHAR(30));

INSERT INTO Hotels (ID, HotelName, HotelInfo, HotelAddr,HotelStar,HotelAmenities,Price,Country) VALUES
(1,"PARKROYAL", "Polished rooms in a chic hotel featuring an outdoor pool, 2 restaurants & a gym", "181 Kitchener Rd, Singapore 208533", 4, "Pool,Parking,WIFI,Air Conditioning",244,"Singapore"),(2,"PARK REGIS SINGAPORE", "Warm lodging in a sleek building with airy quarters & 2 restaurants, plus an outdoor pool & a gym.", "23 Merchant Rd, Singapore 058268", 4, "Pool,Breakfast,WIFI,Air Conditioning",215,"Singapore"),(3,"First World Hotel", "Modest rooms, some with mountain views, in a massive budget hotel set in an entertainment resort", "Genting Highlands, 69000 Genting Highlands, Pahang, Malaysia", 3, "Spa,Fitness Centre,WIFI,Air Conditioning",30,"Malaysia");


INSERT INTO Hotels (ID, HotelName, HotelInfo, HotelAddr,HotelStar,HotelAmenities,Price,Country) VALUES
(4,"Lion Peak Hotel Raffles", "Hip rooms in a cool hotel featuring a restaurant & a rooftop garden, plus complimentary Wi-Fi.", "23 Middle Rd, Singapore 188933", 3, "Breakfast,Parking,WIFI,Air Conditioning",103,"Singapore"),(5,"Park Avenue Rochester", "Chic quarters, some with kitchens, in a casual hotel with a gym, a hot tub & an outdoor pool.", "31 Rochester Dr, Singapore 138637", 4, "Pool,Fitness Centre,WIFI,Air Conditioning",147,"Singapore"),(6,"The Prestige Hotel", "Chic quarters in a high-end hotel offering an elegant restaurant, a gym & a rooftop infinity pool.", "8, Gat Lebuh Gereja, George Town, 10300 George Town, Pulau Pinang, Malaysia", 5, "Rooftop pool,Fitness Centre,WIFI,Air Conditioning",138,"Malaysia");


INSERT INTO Hotels (ID, HotelName, HotelInfo, HotelAddr,HotelStar,HotelAmenities,Price,Country) VALUES
(7,"Al Meroz Hotel", "Warmly decorated rooms in an Arabian-inspired hotel offering a rooftop pool & halal dining.", "4 Ramkhamhaeng 5 Alley, Suan Luang, Bangkok 10250, Thailand", 4, "Pool,Spa,WIFI,Fitness Centre",62,"Thailand"),(8,"Bali Garden Beach Resort", "Rooms & villas in an upscale resort with restaurants & bars, plus 3 pools, a spa & beach access.", "Jl. Kartika Plaza, Kuta, Kec. Kuta, Kabupaten Badung, Bali 80361, Indonesia", 4, "Pool,Spa,Breakfast,Beachfront,WIFI,Air Conditioning",94,"Indonesia"),(9,"Hilton Adelaide", "Warm rooms & suites in a modern high-rise lodging featuring chic dining, a lively bar & a pool.", "233 Victoria Square, Adelaide SA 5000, Australia", 5, "Pool,Breakfast,WIFI,Air Conditioning",361,"Australia");

