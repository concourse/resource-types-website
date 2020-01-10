module Terms.TermsTests exposing (suite)

import Expect exposing (equal)
import Terms.Styles as Styles
import Terms.Terms as Terms
import Test exposing (Test, describe, test)


suite : Test
suite =
    describe "Terms Component"
        [ describe "container"
            [ test "has a width" <|
                \_ ->
                    termsContainer.width |> Expect.equal Styles.containerWidth
            , test "has spacing" <|
                \_ ->
                    termsContainer.spacing |> Expect.equal Styles.containerSpacing
            ]
        , describe "title"
            [ test "has a font size" <|
                \_ ->
                    termsTitle.size |> Expect.equal Styles.titleSize
            , test "has a font" <|
                \_ ->
                    termsTitle.font |> Expect.equal Styles.titleFont
            ]
        , describe "body"
            [ test "has a font size" <|
                \_ ->
                    termsBody.size |> Expect.equal Styles.bodySize
            , test "has a font" <|
                \_ ->
                    termsBody.font |> Expect.equal Styles.bodyFont
            ]
        ]


termsContainer =
    Terms.container


termsTitle =
    Terms.title


termsBody =
    Terms.body
