module Banner.View exposing (view)

import Banner.Banner exposing (banner)
import Common.Common as Common exposing (center)
import Element exposing (Element, column, fill, fromRgb255, height, padding, paragraph, px, text, width)
import Element.Background as Background exposing (color, image)
import Element.Font as Font exposing (center, color, family, size, typeface)


view : Element msg
view =
    let
        container =
            banner.container

        title =
            banner.title

        body =
            banner.body
    in
    column
        [ height <| px container.height
        , width fill
        , Background.color <| fromRgb255 container.backgroundColor
        , Background.image container.backgroundImage
        ]
        [ paragraph
            [ padding title.lineHeight
            , Font.family [ Font.typeface title.font ]
            , Font.color <| fromRgb255 title.color
            , Font.size title.size
            , Font.center
            ]
            [ text title.text ]
        , paragraph
            [ Font.family [ Font.typeface body.font ]
            , Font.color <| fromRgb255 body.color
            , Font.size body.size
            , Font.center
            , width <| px body.width
            , Common.center
            ]
            [ text body.text ]
        ]
