module Card.View exposing (view)

import Card.Card exposing (card)
import Common.Overrides as Overrides exposing (ellipsis)
import Element
    exposing
        ( Color
        , Element
        , clip
        , clipY
        , column
        , el
        , fill
        , fromRgb255
        , height
        , html
        , image
        , maximum
        , mouseOver
        , paddingEach
        , paragraph
        , px
        , rgba255
        , spacing
        , text
        , width
        )
import Element.Border exposing (rounded, shadow)
import Element.Font as Font exposing (color, family, size, typeface)
import Html


edges : { top : Int, right : Int, bottom : Int, left : Int }
edges =
    { top = 0
    , right = 0
    , bottom = 0
    , left = 0
    }


view : Element msg
view =
    let
        container =
            card.container

        name =
            card.resourceType.name

        description =
            card.resourceType.description

        github =
            card.resourceType.github
    in
    el
        [ width <| px container.width
        , height <| px container.height
        , rounded container.borderRadius
        , paddingEach { edges | left = container.paddingLeft }
        , mouseOver
            [ shadow hoverShadow ]
        , shadow cardShadow
        ]
        (column
            []
            [ paragraph
                [ Font.size name.size
                , Font.family [ Font.typeface name.font ]
                , Font.color <| fromRgb255 name.color
                , width (fill |> maximum name.maxWidth)
                , height (fill |> maximum name.maxHeight)
                , paddingEach { edges | top = name.paddingTop }
                , clip
                ]
                -- TODO: unfortunate way of ellipsis
                [ html
                    (Html.div
                        Overrides.ellipsis
                        [ Html.text "Pager Duty" ]
                    )
                ]
            , paragraph
                [ Font.size description.size
                , Font.family [ Font.typeface description.font ]
                , Font.color <| fromRgb255 description.color
                , width (fill |> maximum description.maxWidth)
                , height (fill |> maximum description.maxHeight)
                , spacing description.spacing
                , paddingEach { edges | top = description.paddingTop }
                , clipY
                ]
                [ text "Sends alerts to Pagerduty. This resource can now send Sends alerts to Pagerduty. This resource can now send Sends alerts to Pagerduty." ]
            , paragraph [ paddingEach { edges | top = github.paddingTop } ]
                [ image
                    [ height <| px github.imageHeight
                    , width <| px github.imageWidth
                    ]
                    { src = github.imageName
                    , description = ""
                    }
                ]
            ]
        )


cardShadow : { offset : ( Float, Float ), blur : Float, size : Float, color : Color }
cardShadow =
    let
        container =
            card.container
    in
    { offset =
        ( container.shadow.offsetX
        , container.shadow.offsetY
        )
    , blur = container.shadow.blur
    , size = container.shadow.size
    , color =
        rgba255 container.shadow.color.red
            container.shadow.color.blue
            container.shadow.color.green
            container.shadow.color.alpha
    }


hoverShadow : { offset : ( Float, Float ), blur : Float, size : Float, color : Color }
hoverShadow =
    let
        container =
            card.container
    in
    { offset =
        ( container.hoverShadow.offsetX
        , container.hoverShadow.offsetY
        )
    , blur = container.hoverShadow.blur
    , size = container.hoverShadow.size
    , color =
        rgba255 container.hoverShadow.color.red
            container.hoverShadow.color.blue
            container.hoverShadow.color.green
            container.hoverShadow.color.alpha
    }
