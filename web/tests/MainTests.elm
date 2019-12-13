module MainTests exposing (suite)

import Common.Common as Common exposing (ResourceType)
import Expect exposing (equal)
import Http
import Json.Decode exposing (decodeString)
import Main exposing (Model, Msg, buildErrorMessage, layout, resourceTypeDecoder, update, view)
import Test exposing (Test, describe, test)


suite : Test
suite =
    describe "Main"
        [ describe "json decoder"
            [ test "properly decodes a resource type from good json" <|
                \_ ->
                    let
                        decodedOutput =
                            Json.Decode.decodeString resourceTypeDecoder goodJson
                    in
                    Expect.equal decodedOutput
                        (Ok
                            { name = "some name"
                            , description = "some description"
                            , url = "http://www.example.com"
                            }
                        )
            , test "errors with missing fields in json" <|
                \_ ->
                    let
                        decodedOutput =
                            Json.Decode.decodeString resourceTypeDecoder missingJson
                    in
                    Expect.err decodedOutput
            , test "errors with invalid json" <|
                \_ ->
                    let
                        decodedOutput =
                            Json.Decode.decodeString resourceTypeDecoder invalidJson
                    in
                    Expect.err decodedOutput
            , test "errors with no json" <|
                \_ ->
                    let
                        decodedOutput =
                            Json.Decode.decodeString resourceTypeDecoder ""
                    in
                    Expect.err decodedOutput
            ]
        , describe "build error message"
            [ test "handles bad url by returning the message" <|
                \_ ->
                    Expect.equal
                        (buildErrorMessage <|
                            Http.BadUrl "oh no"
                        )
                        "oh no"
            , test "handles timeouts with a relevant message" <|
                \_ ->
                    Expect.equal True <|
                        String.contains "Server is taking too long to respond." (buildErrorMessage <| Http.Timeout)
            , test "handles network errors with a relevant message" <|
                \_ ->
                    Expect.equal True <|
                        String.contains "Unable to reach server." (buildErrorMessage <| Http.NetworkError)
            , test "handles bad status errors by returning the status" <|
                \_ ->
                    Expect.equal True <|
                        String.contains "500" (buildErrorMessage <| Http.BadStatus 500)
            , test "handles bad body errors by returning the message" <|
                \_ ->
                    Expect.equal
                        (buildErrorMessage <|
                            Http.BadBody "oh no"
                        )
                        "oh no"
            ]
        , describe "update function"
            [ test "updates model with resource types with successful api request" <|
                \_ ->
                    let
                        resourceType =
                            { name = "hi", url = "http", description = "asdf" }

                        model =
                            { resourceTypes = [ resourceType ], errorMessage = Nothing }
                    in
                    Expect.equal (update DataReceived model) ( { resourceTypes = [], errorMessage = Nothing }, Cmd.none )
            ]
        ]


goodJson =
    """
    { "name" : "some name",
    "description" : "some description",
    "url" : "http://www.example.com" }
"""


missingJson =
    """
    { "name" : "some name"}
"""


invalidJson =
    """
    {blah}
"""



-- to test:
-- sets error message when getting a bad http status code?
-- differentiate between different errors?
-- parses json properly and throws it in ye old model?
-- displays cards if proper data?
-- displays error if bad response?
