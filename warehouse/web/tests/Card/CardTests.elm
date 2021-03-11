module Card.CardTests exposing (suite)

import Card.Styles as Styles
import Card.View as Card exposing (view)
import Element exposing (layout)
import Expect exposing (equal)
import Html
import Test exposing (Test, describe, test)
import Test.Html.Query as Query
import Test.Html.Selector exposing (class, tag, text)


suite : Test
suite =
    describe "Card Component"
        [ describe "Github Host"
            [ test "has a github logo" <|
                \_ ->
                    githubHostCard
                        |> Query.fromHtml
                        |> Query.find [ tag "a" ]
                        |> Query.find [ class "host" ]
                        |> Query.has [ class "github-image" ]
            , test "shows number of stars" <|
                \_ ->
                    githubHostCard
                        |> Query.fromHtml
                        |> Query.find [ tag "a" ]
                        |> Query.find [ class "host" ]
                        |> Query.has [ text "123" ]
            , test "shows author" <|
                \_ ->
                    githubHostCard
                        |> Query.fromHtml
                        |> Query.find [ tag "a" ]
                        |> Query.has [ text "author" ]
            ]
        , describe "Non github host"
            [ test "has a different logo logo" <|
                \_ ->
                    otherHostCard
                        |> Query.fromHtml
                        |> Query.find [ tag "a" ]
                        |> Query.find [ class "host" ]
                        |> Query.has [ class "other-image" ]
            , test "does not have github stars" <|
                \_ ->
                    otherHostCard
                        |> Query.fromHtml
                        |> Query.find [ tag "a" ]
                        |> Query.hasNot [ text "123" ]
            , test "does not have show author" <|
                \_ ->
                    otherHostCard
                        |> Query.fromHtml
                        |> Query.find [ tag "a" ]
                        |> Query.hasNot [ text "author" ]
            ]
        ]


githubHostCard =
    Element.layout [] (view <| resourceType "github")


otherHostCard =
    Element.layout [] (view <| resourceType "other")


resourceType host =
    { name = "test"
    , url = "http://example.com"
    , description = "something"
    , username = "author"
    , stars = "123"
    , host = host
    }


view rt =
    Card.view rt
        "http://example.com/example.jpg"
        "http://example.com/example.jpg"
        "http://example.com/example.jpg"
