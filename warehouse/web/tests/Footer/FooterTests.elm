module Footer.FooterTests exposing (..)

import Expect exposing (equal)
import Footer.Footer as Footer
import Footer.Styles as Styles
import Test exposing (Test, describe, test)


suite : Test
suite =
    describe "Footer component"
        [ describe "container"
            [ test "has height" <|
                \_ ->
                    footerContainer.height |> Expect.equal Styles.footerHeight
            , test "has background color" <|
                \_ ->
                    footerContainer.backgroundColor |> Expect.equal Styles.footerBackgroundColor
            , test "has spacing for content" <|
                \_ ->
                    footerContainer.spacing |> Expect.equal Styles.footerContentSpacing
            ]
        , describe "link"
            [ test "has a font" <|
                \_ ->
                    footerLink.font |> Expect.equal Styles.linkFont
            , test "has a font size" <|
                \_ ->
                    footerLink.size |> Expect.equal Styles.linkSize
            , test "has a font color" <|
                \_ ->
                    footerLink.color |> Expect.equal Styles.linkColor
            ]
        ]


footerContainer =
    Footer.container


footerLink =
    Footer.link
