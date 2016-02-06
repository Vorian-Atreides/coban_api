--
-- Fulfill companies
--

INSERT INTO companies(name) VALUES
	  ("アコム株式会社"),
    ("株式会社愛知銀行"),
    ("AOCホールディングス株式会社"),
    ("旭化成株式会社"),
    ("株式会社バンダイ");

--
-- Fulfill addresses
--

INSERT INTO addresses(zip, street, city, company_id) VALUES
	  ("100-8307", "千代田区丸の内二丁目1番1号明治安田生命ビル", "東京都", 1),
    ("23106-1", "中区栄3-14-12", "名古屋市", 2),
    ("140-0002", "品川区東品川二丁目5番8号", "東京都", 3),
    ("101-8101", "千代田区神田神保町1丁目105番地", "東京都", 4),
	  ("111-8081", "台東区駒形1丁目4-8", "東京都", 5),
    ("111-8081", "東京都台東区駒形2-5-4", "東京都", 5);

--
-- Fulfill addresses
--

INSERT INTO accounts(email, password, scope) VALUES
	  ("user@coban.jp", "b14361404c078ffd549c03db443c3fede2f3e534d73f78f77301ed97d4a436a9fd9db05ee8b325c0ad36438b43fec8510c204fc1c1edb21d0941c00e9e2c1ce2", 1), 	## user
    ("office@coban.jp", "f358a8caf95e1889d88444d054c847506d1448fcf03336a621bb9d62ad228d47ca467c0a61d56933f59cc59edb0688270549b4c5d17a6f4937b077d643b868ce", 2),	## office
    ("admin@coban.jp", "c7ad44cbad762a5da0a452f9e854fdc1e0e7a52a38015f23f3eab1d80b931dd472634dfac71cd34ebc35d16ab7fb8a90c81f975113d6c7538dc69dd8de9077ec", 4),	## admin
    ("root@coban.jp", "99adc231b045331e514a516b4b7680f588e3823213abe901738bc3ad67b2f6fcb3c64efb93d18002588d3ccc1a49efbae1ce20cb43df36b38651f11fa75678e8", 7),   ## root
    ("other@coban.jp", "e25ac3845f8cbe12801a2dfa5a89d4c55dc47900f3b6edc9a9ee590f3c2b9312f665d0039c93828b7b58f33950bc817a0955a9c5000a8d3e280569f08745ca68", 1), 	## other
    ("other2@coban.jp", "acd97a6214b6648dd2859dcc48afe1f2ad0603a634d60652f00028f17df150d4298337101cee4456ffa326f18077f2cc1f6ba9f52d13020816127fa0e4378fdf", 1); ## other2


--
-- Fulfill users
--

INSERT INTO users(first_name, last_name, account_id, company_id) VALUES
	  ("青木", "真琳", 1, 1),
    ("織田", "信長", 2, 2),
    ("豊臣", "秀吉", 3, 3),
    ("徳川", "家康", 4, 4);

--
-- Fulfill devices
--

INSERT INTO devices(is_paired, user_id) VALUES
    (false, 1),
    (true, 4);

--
-- Fulfill stations
--

INSERT INTO stations(name, type) VALUES
	  ("銀座線", "metro"),
    ("日比谷線", "metro"),
    ("千代田線", "metro"),
    ("南北線", "metro"),
    ("横須賀駅", "train"),
    ("大宮", "train");

--
-- Fulfill transport histories
--

INSERT INTO transport_histories(date, stock, expense, entrance_id, exit_id, user_id) VALUES
	  (STR_TO_DATE('2016-01-10 06:30:00', '%Y-%m-%d %H:%i:%s'), 850, 150, 1, 2, 1),
    (STR_TO_DATE('2016-01-10 14:10:00', '%Y-%m-%d %H:%i:%s'), 800, 50, 2, 3, 1),
    (STR_TO_DATE('2016-01-10 22:45:00', '%Y-%m-%d %H:%i:%s'), 600, 200, 3, 6, 1),
    (STR_TO_DATE('2016-02-06 04:30:00', '%Y-%m-%d %H:%i:%s'), 10000, 500, 5, 6, 4),
    (STR_TO_DATE('2016-02-06 12:25:00', '%Y-%m-%d %H:%i:%s'), 8000, 2000, 6, 1, 4),
    (STR_TO_DATE('2016-02-06 18:55:00', '%Y-%m-%d %H:%i:%s'), 7500, 500, 1, 3, 4);