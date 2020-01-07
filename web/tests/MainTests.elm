module MainTests exposing (suite)

import Common.Common as Common exposing (ResourceType)
import Expect exposing (equal)
import Http
import Json.Decode exposing (decodeString)
import Main exposing (Model, Msg(..), Page(..), buildErrorMessage, resourceTypeDecoder, update, view)
import RemoteData exposing (RemoteData, WebData)
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
            , test "properly decodes a resource type when description is missing in json" <|
                \_ ->
                    let
                        decodedOutput =
                            Json.Decode.decodeString resourceTypeDecoder missingDescriptionJson
                    in
                    Expect.equal decodedOutput
                        (Ok
                            { name = "some name"
                            , description = ""
                            , url = "http://www.example.com"
                            }
                        )
            , test "errors when name is missing in json" <|
                \_ ->
                    let
                        decodedOutput =
                            Json.Decode.decodeString resourceTypeDecoder missingNameJson
                    in
                    Expect.err decodedOutput
            , test "errors when url is missing in json" <|
                \_ ->
                    let
                        decodedOutput =
                            Json.Decode.decodeString resourceTypeDecoder missingUrlJson
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
                        (buildErrorMessage <| Http.BadUrl "oh no")
                        "oh no"
            , test "handles timeouts with a relevant message" <|
                \_ ->
                    Expect.equal True <|
                        String.contains "Server is taking too long to respond."
                            (buildErrorMessage <| Http.Timeout)
            , test "handles network errors with a relevant message" <|
                \_ ->
                    Expect.equal True <|
                        String.contains "Unable to reach server."
                            (buildErrorMessage <| Http.NetworkError)
            , test "handles bad status errors by returning the status" <|
                \_ ->
                    Expect.equal True <|
                        String.contains "500"
                            (buildErrorMessage <| Http.BadStatus 500)
            , test "handles bad body errors by returning the message" <|
                \_ ->
                    Expect.equal
                        (buildErrorMessage <| Http.BadBody "oh no")
                        "oh no"
            ]
        ]


goodJson =
    """
    { "name" : "some name",
    "description" : "some description",
    "repo" : "http://www.example.com" }
"""


missingNameJson =
    """
    { "description" : "some description",
    "repo" : "http://www.example.com"}
"""


missingDescriptionJson =
    """
    { "name" : "some name",
    "repo" : "http://www.example.com"}
"""


missingUrlJson =
    """
    { "name" : "some name",
    "description": "some description"}
"""


invalidJson =
    """
    {blah}
"""
