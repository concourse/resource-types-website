module Footer.View exposing (view)

import Common.Common as Common exposing (center)
import Element exposing (Element, alignBottom, centerX, column, el, fill, fromRgb255, height, padding, paragraph, px, row, text, width)
import Element.Background as Background exposing (color, image)
import Element.Font as Font exposing (center, color, family, size, typeface)
import Footer.Footer exposing (footer)


view : Element msg
view =
    let
        container =
            footer.container
    in
    row
        [ height <| px container.height
        , width <| fill
        , Background.color <| fromRgb255 container.backgroundColor
        , Font.color <| fromRgb255 container.color
        , Font.size container.size
        , Font.family [ Font.typeface container.font ]
        , Common.center
        ]
        [ el [ centerX ] (text "Terms of Use | Contribute | Feedback") ]



-- TODO:
-- get the links in the footer object as an array? styles and text? or hardcode? :SHRUG:
-- need to get the acutal links
-- need to do the router
-- figure out the asset thing
