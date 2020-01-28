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
                    Terms.containerWidth |> Expect.equal Styles.containerWidth
            ]
        , describe "back link"
            [ test "has a font size" <|
                \_ -> termsBackLink.size |> Expect.equal Styles.backLinkSize
            , test "has top padding" <|
                \_ -> termsBackLink.paddingTop |> Expect.equal Styles.backLinkPaddingTop
            , test "has font color" <|
                \_ -> termsBackLink.color |> Expect.equal Styles.backLinkColor
            ]
        , describe "title"
            [ test "has a font size" <|
                \_ ->
                    termsTitle.size |> Expect.equal Styles.titleSize
            , test "has a font" <|
                \_ ->
                    termsTitle.font |> Expect.equal Styles.titleFont
            , test "has padding" <|
                \_ ->
                    termsTitle.padding |> Expect.equal Styles.titlePadding
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


termsBackLink =
    Terms.backLink


termsTitle =
    Terms.title


termsBody =
    Terms.body
