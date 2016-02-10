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
