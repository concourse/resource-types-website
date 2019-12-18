module Footer.FooterTests exposing (..)

import Expect exposing (equal)
import Footer.Footer as Footer
import Footer.Styles as Styles
import Test exposing (Test, describe, test)


suite : Test
suite =
    describe "Footer component"
        [ test "has height" <|
            \_ ->
                footerContainer.height |> Expect.equal Styles.footerHeight
        , test "has background color" <|
            \_ ->
                footerContainer.backgroundColor |> Expect.equal Styles.footerBackgroundColor
        , test "has a font" <|
            \_ ->
                footerContainer.font |> Expect.equal Styles.footerFont
        , test "has a font size" <|
            \_ -> footerContainer.size |> Expect.equal Styles.footerSize
        , test "has a font color" <|
            \_ -> footerContainer.color |> Expect.equal Styles.footerColor
        ]


footerContainer =
    Footer.container
