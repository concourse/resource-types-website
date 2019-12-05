module Card.CardTests exposing (..)

import Card.Card as Card exposing (container)
import Card.Styles as Styles exposing (..)
import Expect exposing (equal)
import Test exposing (Test, describe, test)


suite : Test
suite =
    describe "Card Component"
        [ describe "container"
            [ test "has a height" <|
                \_ ->
                    cardContainer.height
                        |> Expect.equal Styles.height
            , test "has a width" <|
                \_ ->
                    cardContainer.width
                        |> Expect.equal Styles.width
            , test "has a border radius" <|
                \_ ->
                    cardContainer.borderRadius
                        |> Expect.equal Styles.borderRadius
            , test "has a shadow" <|
                \_ ->
                    cardContainer.shadow
                        |> Expect.equal Styles.shadow
            , test "has a hover shadow" <|
                \_ ->
                    cardContainer.hoverShadow
                        |> Expect.equal Styles.hoverShadow
            , test "has spacing" <|
                \_ ->
                    cardContainer.spacing
                        |> Expect.equal Styles.spacing
            ]
        ]


cardContainer =
    Card.container
