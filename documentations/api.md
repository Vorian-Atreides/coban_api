FORMAT: 1A
# ATES

This API is allowing the consumers to interact with the data related to the Coban project.

# Group Authentications

Actions related to the authentication

## Administration's authentication [/administrations/authenticate]

### Authenticate [POST]

+ Request

        {
            "email":"admin@coban.jp",
            "password":"admin"
        }

+ Response 200

        eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjEuNDU1MDUzNzA3ZSswOSwiaWF0IjoxLjQ1NTA0NjUwN2UrMDksImlzcyI6ImNvYmFuIiwibmJmIjoxLjQ1NTA0NjUwN2UrMDksInNjb3BlIjo0LCJ1c2VyIjozfQ.SXE5VMg2Hz_zwNB35VFqF1kXzLMpKXKSzuduqUPH_qfe9WCqkJWktyj5rr-6-R3MnfnNc45vKQW7dfQDGw5zK1x102G9c9gjefmJsN0rmHrbeYnHgODsDrxQ20910DOKcFouhz-_BES_lXRbVdYk0c2-OqhrcuxLxpHwEDzvpx6FWORwBAU6duCOKspoeI4O1VoJmaL8BirXqTH5OJWVgodeBtKIYyvK2PPBlKE-1Xgvuc5TcQBtfIz8I738jENirD9UnkSwOL7H9GtZZmGW0V1cjMUmL1aLbSoUFhmiqrtwzH-AQuF1GQvOW8hXIDdVR-PO2u8sCL_SjK8BPqrlVkGROzIC-AdTKHq8kDWRNx5zeQhr3Jr-bkSEx-XrfhFOHSHq_RAppvDg3c-8cr3Ss8dNnc9WrmFL-f23T0LLGUizBekGkevXnuXBB6r90NFhPbdIkQTuWH4BUmkFDhPLuo8mk5mwP9LGo7S28Dm6IerpnXJ3bfL7_vbhA1Zwwco-xONfMYSGWNqhptjSBNmV5wQQVtvLMe3EpVOBUpJ7K-umXNFWM4jqJUJhSBJjMhWcu2hCM_tlo9Q3kKLtBLRhtRQKg898MokiB93hTCqBCXaUrEnjbFNPN8HZgBSAmsL0Invk-hIh2XWMkdNh0e5LG1IAyZv-KGPEBHljG02GmDY

+ Response 400

        Invalid JSON.

+ Response 401

        The credentials are invalid.

## Office's authentication [/offices/authenticate]

### Authenticate [POST]

+ Request

        {
            "email":"office@coban.jp",
            "password":"office"
        }

+ Response 200

        eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjEuNDU1MDU0NzVlKzA5LCJpYXQiOjEuNDU1MDQ3NTVlKzA5LCJpc3MiOiJjb2JhbiIsIm5iZiI6MS40NTUwNDc1NWUrMDksInNjb3BlIjoyLCJ1c2VyIjoyfQ.357NdG4mWIdmH-NB0_jNJPmiO8RBxLHvwwSyXnzfnUXw8zvlh83ndeCRqhn8G2_5hOulPAOVWQGg_Lhq0bnb3QerFtCF0rIX3ju0-Cf3cGtfbH6ConMDr28N4lPzUOTl-CMLqO2ctODxyL-6n3T-OxX9dtF-vqeBdypCoCFr51BgRafElA5G_i1FFCYwxht9BQhbfygyy8-9DucEnCVY5wGrMrBJIgFUJnOiSFaPJ9HYAfs017Jeyf3UjSPI6pOZfkLRvZ09xRvPGLRMKNYf2Pn7-5hNSPkRSAVeVClS5XbS44dd9t5NWuYTbMZck56hYZo_KrO5mcYxoJyiHajd88bxlgGu1VVj6D_5Ss05fVzD3bXPiKVXsOfaAcAKdKnFkJryELilCq3lszxUPVsRl1AxWVpqEC4yN38GJCv7Qi9niSdKnbvb-hgVZpNsNh8T0Crp-w7gsFY42dbd1086Xl4VVEa6aMSfLxrLUHya-Qzu1xXY_KZW39EFk4zYo0QmVu2-cqcxcMGOxAJ623X2KmNT_-F5D8w0aHXbmOwmoEjS1pcfh7L7prlA0jB4Z6m_wl9nSqmgBY5BQGYLIRwoNlXLrYqTUG2s0ycC6fl_6cDsqjUJC32tATbIqpBkLTh0oTUXldVBaOKwZ_XRSVedC9bXpk94LV2CEBgjpZ2DZ2Y


+ Response 400

        Invalid JSON.

+ Response 401

        The credentials are invalid.

## Client's authentication [/clients/authenticate]

### Authenticate [POST]

+ Request

        {
            "email":"user@coban.jp",
            "password":"user"
        }

+ Response 200

        eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjEuNDU1MDU0ODQ1ZSswOSwiaWF0IjoxLjQ1NTA0NzY0NWUrMDksImlzcyI6ImNvYmFuIiwibmJmIjoxLjQ1NTA0NzY0NWUrMDksInNjb3BlIjoxLCJ1c2VyIjoxfQ.HH2fAD0mpdA9WVv8c9mmkfuaHjkjGQ6S2GfUA07sLe9WhMZvS4nsjf_U8blfuXTM8FgJ3oQ6QwuGcHTePRGL6NR3Me9a9cNX6Ds-EIOWIWhYJiEfcCOi8Slm5MsvpAOTXmnRqfxg4WzGbr-2A8Rz9t_8UHsX3HRyTIZ7HzU-tMTCmaGNSIxQyxOBx0MmYGd7pWA8h8N7-ERlI5kvIBOvgQsRcYb6xNtrSkPdWlQtqgy4stzenBd0iPBsMOLEW3iw_8exnOwx_-BJVBVhSD51pws1D0rfqJ2xxcKWJWKYCeLwSriKGgylckGSdMevCZthqVMnGnrLTp94w3n7-wr-Fsj_AdnFl8JoV7qo6xO9BtAHuQcbFrSmLRAdLXaLBm7SU4SDhGBO3mSjDaoIJWgeiThakeaMSXLNW0E2cS1ifGfysy2Vh5Oa5Sf_geocpX9gF98vPQfbUY55V9PzXWektvm7vEfw93pfJ_gXPV0yrZMnjf5KILqbQdWpBjutsb_yXlybqs7f7V0kRo5qGJ8q0zee6pIBKzuOjUUqrwkujGS6KMBMp9M7eeYfVYpkwpvUYPNFhcdjEwO1TP_MT4E_4F3aqPNLLCCjjge-V-dD7nzNjlcDMdJIAsXohxgvdZj5dy5bzvU-ezP9lOrN091GmHIXpJoVtmxozcRb7g1tXKw

+ Response 400

        Invalid JSON.

+ Response 401

        The credentials are invalid.

# Group Addresses

Resources related to the addresses.

## Administration Addresses [/administrations/addresses]

### Every addresses [GET]

+ Request

    + Headers

            Authorization: Bearer <token>

+ Response 200 (application/json)

        [
            {
                "id":1,
                "zip":"100-8307",
                "street":"千代田区丸の内二丁目1番1号明治安田生命ビル",
                "city":"東京都",
                "company-id":1
            },
            {
                "id":2,
                "zip":"23106-1",
                "street":"中区栄3-14-12",
                "city":"名古屋市",
                "company-id":2
            },
            {
                "id":3,
                "zip":"140-0002",
                "street":"品川区東品川二丁目5番8号",
                "city":"東京都",
                "company-id":3
            },
            {
                "id":4,
                "zip":"101-8101",
                "street":"千代田区神田神保町1丁目105番地",
                "city":"東京都",
                "company-id":4
            },
            {
                "id":5,
                "zip":"111-8081",
                "street":"台東区駒形1丁目4-8",
                "city":"東京都",
                "company-id":5
            },
            {
                "id":6,
                "zip":"111-8081",
                "street":"東京都台東区駒形2-5-4",
                "city":"東京都",
                "company-id":5
            }
        ]

+ Response 400

        The credentials are invalid.

## Office Addresses [/offices/addresses]

### Company's addresses [GET]

+ Request

    + Headers

            Authorization: Bearer <token>

+ Response 200 (application/json)

        [
            {
                "id":5,
                "zip":"111-8081",
                "street":"台東区駒形1丁目4-8",
                "city":"東京都",
                "company-id":5
            },
            {
                "id":6,
                "zip":"111-8081",
                "street":"東京都台東区駒形2-5-4",
                "city":"東京都",
                "company-id":5
            }
        ]

+ Response 400

        The credentials are invalid.

### Add an addresses [POST]

+ Request

    + Headers

            Authorization: Bearer <token>

    + Body

            {
                "zip":"222-222",
                "street":"25 Ginza",
                "city":"Tokyo"
            }

+ Response 201 (application/json)

        {
            "id":7,
            "zip":"222-222",
            "street":"25 Ginza",
            "city":"Tokyo",
            "company-id":4
        }

+ Response 400

        The credentials are invalid.

# Group Accounts

Resources related to the accounts.

## Administration Accounts [/administrations/accounts]

### Every accounts [GET]

+ Request

    + Headers

            Authorization: Bearer <token>

+ Response 200 (application/json)

        [
            {
                "email":"user@coban.jp"
            },
            {
                "email":"office@coban.jp"
            },
            {
                "email":"admin@coban.jp"
            },
            {
                "email":"root@coban.jp"
            }
        ]

+ Response 400

        The credentials are invalid.

# Group Users

Resources related to the users.

## Administration Users [/administrations/users]

### Every users [GET]

+ Request

    + Headers

            Authorization: Bearer <token>

+ Response 200 (application/json)

        [
            {
                "id":1,
                "first-name":"青木",
                "last-name":"真琳",
                "Account":{
                    "email":"user@coban.jp"
                },
                "Company":{
                    "id":1,
                    "name":"アコム株式会社"
                },
                "Device":{
                    "is-paired":false
                }
            },
            {
                "id":2,
                "first-name":"織田",
                "last-name":"信長",
                "Account":{
                    "email":"office@coban.jp"
                },
                "Company":{
                    "id":2,
                    "name":"株式会社愛知銀行"
                },
                "Device":{
                    "is-paired":false
                }
            },
            {
                "id":3,
                "first-name":"豊臣",
                "last-name":"秀吉",
                "Account":{
                    "email":"admin@coban.jp"
                },
                "Company":{
                    "id":3,
                    "name":"AOCホールディングス株式会社"
                },
                "Device":{
                    "is-paired":false
                }
            },
            {
                "id":4,
                "first-name":"徳川",
                "last-name":"家康",
                "Account":{
                    "email":"root@coban.jp"
                },
                "Company":{
                    "id":4,
                    "name":"旭化成株式会社"
                },
                "Device":{
                    "is-paired":true
                }
            }
        ]

+ Response 400

        The credentials are invalid.

## Office Users [/offices/users]

### Company's employees [GET]

+ Request

    + Headers

            Authorization: Bearer <token>

+ Response 200

        [
            {
                "id":2,
                "first-name":"織田",
                "last-name":"信長",
                "Account":null,
                "Company":null,
                "Device":null
            }
        ]

+ Response 400

        The credentials are invalid.

### Create an employee [POST]

+ scope (string) - Can be "Office" or "Client"

+ Request

    + Headers

            Authorization: Bearer <token>

    + Body

            {
                "first-name":"Gaston",
                "last-name":"Siffert",
                "email":"gs060292@live.fr",
                "scope":"Office"
            }

+ Response 201

        {
            "first-name":"Gaston",
            "last-name":"Siffert",
            "account":null,
            "company":null,
            "device":null
        }

+ Response 400

        The credentials are invalid.

## Office Users [/offices/users/{id}]

### Update an employee [PUT]

+ scope (string) - Can be "Office" or "Client"

+ Parameters
    + id (number) - ID of the user to modify

+ Request

    + Headers

            Authorization: Bearer <token>

    + Body

            {
                "first-name":"Gaston",
                "last-name":"Siffert",
                "email":"gs060292@live.fr",
                "scope":"Office"
            }

+ Response 200

        {
            "first-name":"Gaston",
            "last-name":"Siffert",
            "account":null,
            "company":null,
            "device":null
        }

+ Response 400

        The credentials are invalid.

## Client Users [/clients/users]

### Get the user [GET]

+ Request

    + Headers

            Authorization: Bearer <token>

+ Response 200

            {
                "first-name":"Gaston",
                "last-name":"Siffert",
                "account":null,
                "company":null,
                "device":null
            }

+ Response 400

        The credentials are invalid.

### Update its password [PUT]

+ Request

    + Headers

            Authorization: Bearer <token>

    + Body

            {
                "old-password":"previous_password",
                "password-1":"new_password",
                "password-2":"new_password"
            }

+ Response 200

            {
                "first-name":"Gaston",
                "last-name":"Siffert",
                "account":null,
                "company":null,
                "device":null
            }

+ Response 400

        The credentials are invalid.

# Group Companies

Resources related to the companies

## Administration Companies [/administrations/companies]

### Every companies [GET]

+ Request

    + Headers

            Authorization: Bearer <token>

+ Response 200

        [
            {
                "id":1,
                "name":"アコム株式会社"
            },
            {
                "id":2,
                "name":"株式会社愛知銀行"
            },
            {
                "id":3,
                "name":"AOCホールディングス株式会社"
            },
            {
                "id":4,
                "name":"旭化成株式会社"
            },
            {
                "id":5,
                "name":"株式会社バンダイ"
            }
        ]

+ Response 400

        The credentials are invalid.

## Office Companies [/offices/companies]

### Current company [GET]

+ Request

    + Headers

            Authorization: Bearer <token>

+ Response 200

        {
            "id":1,
            "name":"アコム株式会社"
        }

+ Response 400

        The credentials are invalid.

### Create company [POST]

+ Request

        {
            "name":"Coban",
            "administrator": {
                "first-name":"Tatsuya",
                "last-name":"Zembutsu",
                "email":"tatsuya.zembutsu@coban.jp",
                "password":"password"
            }
        }

+ Response 201

        {
            "id":5,
            "name":"Coban"
        }

## Client Companies [/clients/companies]

### Current company [GET]

+ Request

    + Headers

            Authorization: Bearer <token>

+ Response 200

        {
            "id":5,
            "name":"Coban"
        }

# Group Stations

Resources related to the stations

## Administration Stations [/administrations/stations]

### Every stations [GET]

+ Request

    + Headers

            Authorization: Bearer <token>

+ Response 200

        [
            {
                "name":"銀座線",
                "type":"metro"
            },
            {
                "name":"日比谷線",
                "type":"metro"
            },
            {
                "name":"千代田線",
                "type":"metro"
            },
            {
                "name":"南北線",
                "type":"metro"
            },
            {
                "name":"横須賀駅",
                "type":"train"
            },
            {
                "name":"大宮",
                "type":"train"
            }
        ]

+ Response 400

        The credentials are invalid.

# Group Transport Histories

Resources related to the transport histories

## Administration Transport Histories [/administrations/transport-histories]

### Every transport histories [GET]

+ Request

    + Headers

            Authorization: Bearer <token>

+ Response 200

        [
            {
                "id":1,
                "date":"2016-01-10T06:30:00Z",
                "stock":850,
                "expense":150,
                "Entrance":{
                    "name":"銀座線",
                    "type":"metro"
                },
                "Exit":{
                    "name":"日比谷線",
                    "type":"metro"
                },
                "User":{
                    "id":1,
                    "first-name":"青木",
                    "last-name":"真琳",
                    "Account":null,
                    "Company":null,
                    "Device":null
                }
            },
            {
                "id":2,
                "date":"2016-01-10T14:10:00Z",
                "stock":800,
                "expense":50,
                "Entrance":{
                    "name":"日比谷線",
                    "type":"metro"
                },
                "Exit":{
                    "name":"千代田線",
                    "type":"metro"
                },
                "User":{
                    "id":1,
                    "first-name":"青木",
                    "last-name":"真琳",
                    "Account":null,
                    "Company":null,
                    "Device":null
                }
            },
            {
                "id":3,
                "date":"2016-01-10T22:45:00Z",
                "stock":600,
                "expense":200,
                "Entrance":{
                    "name":"千代田線",
                    "type":"metro"
                },
                "Exit":{
                    "name":"大宮",
                    "type":"train"
                },
                "User":{
                    "id":1,
                    "first-name":"青木",
                    "last-name":"真琳",
                    "Account":null,
                    "Company":null,
                    "Device":null
                }
            },
            {
                "id":4,
                "date":"2016-02-06T04:30:00Z",
                "stock":10000,
                "expense":500,
                "Entrance":{
                    "name":"横須賀駅",
                    "type":"train"
                },
                "Exit":{
                    "name":"大宮",
                    "type":"train"
                },
                "User":{
                    "id":4,
                    "first-name":"徳川",
                    "last-name":"家康",
                    "Account":null,
                    "Company":null,
                    "Device":null
                }
            },
            {
                "id":5,
                "date":"2016-02-06T12:25:00Z",
                "stock":8000,
                "expense":2000,
                "Entrance":{
                    "name":"大宮",
                    "type":"train"
                },
                "Exit":{
                    "name":"銀座線",
                    "type":"metro"
                },
                "User":{
                    "id":4,
                    "first-name":"徳川",
                    "last-name":"家康",
                    "Account":null,
                    "Company":null,
                    "Device":null
                }
            },
            {
                "id":6,
                "date":"2016-02-06T18:55:00Z",
                "stock":7500,
                "expense":500,
                "Entrance":{
                    "name":"銀座線",
                    "type":"metro"
                },
                "Exit":{
                    "name":"千代田線",
                    "type":"metro"
                },
                "User":{
                    "id":4,
                    "first-name":"徳川",
                    "last-name":"家康",
                    "Account":null,
                    "Company":null,
                    "Device":null
                }
            }
        ]

+ Response 400

        The credentials are invalid.

## Office Transport Histories [/offices/transport-histories]

### Employees' transport histories [GET]

+ Request

    + Headers

            Authorization: Bearer <token>

+ Response 200

        [
            {
                "id":4,
                "date":"2016-02-06T04:30:00Z",
                "stock":10000,
                "expense":500,
                "Entrance":{
                    "name":"横須賀駅",
                    "type":"train"
                },
                "Exit":{
                    "name":"大宮",
                    "type":"train"
                },
                "User":{
                    "id":4,
                    "first-name":"徳川",
                    "last-name":"家康",
                    "Account":null,
                    "Company":null,
                    "Device":null
                }
            },
            {
                "id":5,
                "date":"2016-02-06T12:25:00Z",
                "stock":8000,
                "expense":2000,
                "Entrance":{
                    "name":"大宮",
                    "type":"train"
                },
                "Exit":{
                    "name":"銀座線",
                    "type":"metro"
                },
                "User":{
                    "id":4,
                    "first-name":"徳川",
                    "last-name":"家康",
                    "Account":null,
                    "Company":null,
                    "Device":null
                }
            },
            {
                "id":6,
                "date":"2016-02-06T18:55:00Z",
                "stock":7500,
                "expense":500,
                "Entrance":{
                    "name":"銀座線",
                    "type":"metro"
                },
                "Exit":{
                    "name":"千代田線",
                    "type":"metro"
                },
                "User":{
                    "id":4,
                    "first-name":"徳川",
                    "last-name":"家康",
                    "Account":null,
                    "Company":null,
                    "Device":null
                }
            }
        ]

+ Response 400

        The credentials are invalid.

## Office Transport Histories by user [/offices/transport-histories/{id}]

### Employee's transport histories [GET]

+ Parameters
    + id (number) - ID of the user to gather

+ Request

    + Headers

            Authorization: Bearer <token>

+ Response 200

        [
            {
                "id":4,
                "date":"2016-02-06T04:30:00Z",
                "stock":10000,
                "expense":500,
                "Entrance":{
                    "name":"横須賀駅",
                    "type":"train"
                },
                "Exit":{
                    "name":"大宮",
                    "type":"train"
                },
                "User":{
                    "id":4,
                    "first-name":"徳川",
                    "last-name":"家康",
                    "Account":null,
                    "Company":null,
                    "Device":null
                }
            },
            {
                "id":5,
                "date":"2016-02-06T12:25:00Z",
                "stock":8000,
                "expense":2000,
                "Entrance":{
                    "name":"大宮",
                    "type":"train"
                },
                "Exit":{
                    "name":"銀座線",
                    "type":"metro"
                },
                "User":{
                    "id":4,
                    "first-name":"徳川",
                    "last-name":"家康",
                    "Account":null,
                    "Company":null,
                    "Device":null
                }
            },
            {
                "id":6,
                "date":"2016-02-06T18:55:00Z",
                "stock":7500,
                "expense":500,
                "Entrance":{
                    "name":"銀座線",
                    "type":"metro"
                },
                "Exit":{
                    "name":"千代田線",
                    "type":"metro"
                },
                "User":{
                    "id":4,
                    "first-name":"徳川",
                    "last-name":"家康",
                    "Account":null,
                    "Company":null,
                    "Device":null
                }
            }
        ]

+ Response 400

        The credentials are invalid.

## Clients Transport Histories [/clients/transport-histories]

### Current user's transport histories [GET]

+ Request

    + Headers

            Authorization: Bearer <token>

+ Response 200

        [
            {
                "id":1,
                "date":"2016-01-10T06:30:00Z",
                "stock":850,
                "expense":150,
                "Entrance":{
                    "name":"銀座線",
                    "type":"metro"
                },
                "Exit":{
                    "name":"日比谷線",
                    "type":"metro"
                }
            },
            {
                "id":2,
                "date":"2016-01-10T14:10:00Z",
                "stock":800,
                "expense":50,
                "Entrance":{
                    "name":"日比谷線",
                    "type":"metro"
                },
                "Exit":{
                    "name":"千代田線",
                    "type":"metro"
                }
            },
            {
                "id":3,
                "date":"2016-01-10T22:45:00Z",
                "stock":600,
                "expense":200,
                "Entrance":{
                    "name":"千代田線",
                    "type":"metro"
                },
                "Exit":{
                    "name":"大宮",
                    "type":"train"
                }
            }
        ]

+ Response 400

        The credentials are invalid.

### Add new transport histories for the current user [POST]

    FgEAAiBK5SvOIScPAAq6AA== is a base64 string of the data:[0x16, 0x01, 0x00, 0x02, 0x20, 0x4A, 0xE5, 0x2B, 0xCE, 0x21, 0x27, 0x0F, 0x00, 0x0A, 0xBA, 0x00]

    FgEAAiBK4z7jMIsHAAqzAA== is a base64 string of the data: [0x16, 0x01, 0x00, 0x02, 0x20, 0x4A, 0xE3, 0x3E, 0xE3, 0x30, 0x8B, 0x07, 0x00, 0x0A, 0xB3, 0x00]

    FgEAAiBKziHSAk4IAAqxAA== is a base64 string of the data:[0x16, 0x01, 0x00, 0x02, 0x20, 0x4A, 0xCE, 0x21, 0xD2, 0x02, 0x4E, 0x08, 0x00, 0x0A, 0xB1, 0x00]

+ Request

    + Headers

            Authorization: Bearer <token>

    + Body

            [
                "FgEAAiBK5SvOIScPAAq6AA==",
                "FgEAAiBK4z7jMIsHAAqzAA==",
                "FgEAAiBKziHSAk4IAAqxAA=="
            ]
        
+ Response 201

        [
            {
                "id":11,
                "date":"2016-02-10T00:00:00Z",
                "stock":3879,
                "expense":0,
                "Entrance":
                {
                    "company":"東京地下鉄",
                    "line":"4号線丸ノ内",
                    "station":"東京"
                },
                "Exit":
                {
                    "company":"東京急行電鉄",
                    "line":"東急東横",
                    "station":"祐天寺"
                },
                "User":
                {
                    "id":1,
                    "first-name":"青木",
                    "last-name":"真琳",
                    "Account":null,
                    "Company":null,
                    "Device":null
                }
            },
            {
                "id":12,
                "date":"2016-02-10T00:00:00Z",
                "stock":1931,
                "expense":0,
                "Entrance":
                {
                    "company":"東京地下鉄",
                    "line":"3号線銀座",
                    "station":"渋谷"
                },
                "Exit":
                {
                    "company":"東京地下鉄",
                    "line":"3号線銀座",
                    "station":"上野広小路"
                },
                "User":
                {
                    "id":1,
                    "first-name":"青木",
                    "last-name":"真琳",
                    "Account":null,
                    "Company":null,
                    "Device":null
                }
            },
            {
                "id":13,
                "date":"2016-02-10T00:00:00Z",
                "stock":2126,
                "expense":0,
                "Entrance":
                {
                    "company":"東京急行電鉄",
                    "line":"東急東横",
                    "station":"祐天寺"
                },
                "Exit":
                {
                    "company":"東京急行電鉄",
                    "line":"東急田園都市",
                    "station":"渋谷"
                },
                "User":
                {
                    "id":1,
                    "first-name":"青木",
                    "last-name":"真琳",
                    "Account":null,
                    "Company":null,
                    "Device":null
                }
            }
        ]
