module Card.View exposing (view)

import Card.Card exposing (Author, Description, Github, Name, card)
import Common.Common exposing (ResourceType)
import Common.Overrides as Overrides exposing (ellipsis, multiLineEllipsis)
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
        , spacing
        , text
        , width
        )
import Element.Border exposing (rounded, shadow)
import Element.Font as Font exposing (color, family, size, typeface)
import Html


padding : { top : Int, right : Int, bottom : Int, left : Int }
padding =
    { top = 0
    , right = 0
    , bottom = 0
    , left = 0
    }


view : ResourceType -> String -> Element msg
view resourceType githubIcon =
    let
        container =
            card.container
    in
    newTabLink []
        { url = resourceType.url
        , label =
            el
                [ width <| px container.width
                , height <| px container.height
                , rounded container.borderRadius
                , paddingEach { padding | left = container.paddingLeft }
                , mouseOver
                    [ shadow cardHoverShadow ]
                , shadow cardShadow
                ]
                (column
                    []
                    [ name resourceType card.resourceType.name
                    , author resourceType card.resourceType.author
                    , description resourceType card.resourceType.description
                    , github card.resourceType.github githubIcon
                    ]
                )
        }


name : ResourceType -> Name -> Element msg
name resourceType styles =
    paragraph
        [ Font.size styles.size
        , Font.family [ Font.typeface styles.font ]
        , Font.color <| fromRgb255 styles.color
        , height fill
        , width (fill |> maximum styles.maxWidth)
        , paddingEach { padding | top = styles.paddingTop }
        , clip
        ]
        [ html
            (Html.div
                Overrides.ellipsis
                [ Html.text resourceType.name ]
            )
        ]


author : ResourceType -> Author -> Element msg
author resourceType styles =
    paragraph
        [ Font.family
            [ Font.typeface styles.font ]
        , Font.size styles.size
        , Font.color <| fromRgb255 styles.color
        , paddingEach { padding | top = styles.paddingTop }
        ]
        -- it'll be resourceType.author or whatever here
        [ text "@jomsie" ]


description : ResourceType -> Description -> Element msg
description resourceType styles =
    paragraph
        [ Font.size styles.size
        , Font.family [ Font.typeface styles.font ]
        , Font.color <| fromRgb255 styles.color
        , width (fill |> maximum styles.maxWidth)
        , height (fill |> minimum styles.minHeight)
        , spacing styles.spacing
        , paddingEach { padding | top = styles.paddingTop }
        , clipY
        ]
        [ html
            (Html.div
                (Overrides.multiLineEllipsis 2)
                [ Html.text resourceType.description ]
            )
        ]


github : Github -> String -> Element msg
github styles githubIconImg =
    paragraph [ paddingEach { padding | top = styles.paddingTop } ]
        [ image
            [ height <| px styles.imageHeight
            , width <| px styles.imageWidth
            ]
            { src = githubIconImg
            , description = ""
            }
        ]


cardShadow : { offset : ( Float, Float ), blur : Float, size : Float, color : Color }
cardShadow =
    let
        shadow =
            card.container.shadow
    in
    { offset = shadow.offset
    , blur = shadow.blur
    , size = shadow.size
    , color = fromRgb255 shadow.color
    }


cardHoverShadow : { offset : ( Float, Float ), blur : Float, size : Float, color : Color }
cardHoverShadow =
    let
        hoverShadow =
            card.container.hoverShadow
    in
    { offset = hoverShadow.offset
    , blur = hoverShadow.blur
    , size = hoverShadow.size
    , color = fromRgb255 hoverShadow.color
    }
