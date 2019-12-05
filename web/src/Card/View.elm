module Card.View exposing (view)

import Card.Card exposing (card)
import Element exposing (Element, el, height, mouseOver, px, rgba255, spacing, text, width)
import Element.Border exposing (rounded, shadow)


view : Element msg
view =
    let
        style =
            card.container
    in
    el
        [ width <| px style.width
        , height <| px style.height
        , rounded style.borderRadius
        , spacing 16
        , mouseOver
            [ shadow
                { offset =
                    ( style.hoverShadow.offsetX
                    , style.hoverShadow.offsetY
                    )
                , blur = style.hoverShadow.blur
                , size = style.hoverShadow.size
                , color =
                    rgba255 style.hoverShadow.color.red
                        style.hoverShadow.color.blue
                        style.hoverShadow.color.green
                        style.hoverShadow.color.alpha
                }
            ]
        , shadow
            { offset =
                ( style.shadow.offsetX
                , style.shadow.offsetY
                )
            , blur = style.shadow.blur
            , size = style.shadow.size
            , color =
                rgba255 style.shadow.color.red
                    style.shadow.color.blue
                    style.shadow.color.green
                    style.shadow.color.alpha
            }
        ]
        (text "")
