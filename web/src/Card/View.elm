module Card.View exposing (view)

import Card.Card exposing (card)
import Common.Common exposing (ResourceType)
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
        , minimum
        , mouseOver
        , newTabLink
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


view : ResourceType -> Element msg
view resourceType =
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
            [ shadow cardHoverShadow ]
        , shadow cardShadow
        ]
        (newTabLink []
            { url = resourceType.url
            , label =
                column
                    []
                    [ paragraph
                        [ Font.size name.size
                        , Font.family [ Font.typeface name.font ]
                        , Font.color <| fromRgb255 name.color
                        , width (fill |> maximum name.maxWidth)
                        , paddingEach { edges | top = name.paddingTop }
                        , clip
                        ]
                        [ html
                            (Html.div
                                Overrides.ellipsis
                                [ Html.text resourceType.name ]
                            )
                        ]
                    , paragraph
                        [ Font.size description.size
                        , Font.family [ Font.typeface description.font ]
                        , Font.color <| fromRgb255 description.color
                        , width (fill |> maximum description.maxWidth)
                        , height
                            (fill
                                |> maximum description.maxHeight
                                |> minimum description.minHeight
                            )
                        , spacing description.spacing
                        , paddingEach { edges | top = description.paddingTop }
                        , clipY
                        ]
                        [ text resourceType.description ]
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
            }
        )


cardShadow : { offset : ( Float, Float ), blur : Float, size : Float, color : Color }
cardShadow =
    let
        shadow =
            card.container.shadow
    in
    { offset =
        ( shadow.offsetX
        , shadow.offsetY
        )
    , blur = shadow.blur
    , size = shadow.size
    , color =
        rgba255 shadow.color.red
            shadow.color.blue
            shadow.color.green
            shadow.color.alpha
    }


cardHoverShadow : { offset : ( Float, Float ), blur : Float, size : Float, color : Color }
cardHoverShadow =
    let
        hoverShadow =
            card.container.hoverShadow
    in
    { offset =
        ( hoverShadow.offsetX
        , hoverShadow.offsetY
        )
    , blur = hoverShadow.blur
    , size = hoverShadow.size
    , color =
        rgba255 hoverShadow.color.red
            hoverShadow.color.blue
            hoverShadow.color.green
            hoverShadow.color.alpha
    }
