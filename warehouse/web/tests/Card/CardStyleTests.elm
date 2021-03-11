module Card.CardStyleTests exposing (cardContainer, cardResourceType, suite)

import Card.Card as Card exposing (container)
import Card.Styles as Styles
import Expect exposing (equal)
import Test exposing (Test, describe, test)


suite : Test
suite =
    describe "Card Component"
        [ describe "container"
            [ test "has a height" <|
                \_ ->
                    cardContainer.height
                        |> Expect.equal Styles.containerHeight
            , test "has a width" <|
                \_ ->
                    cardContainer.width
                        |> Expect.equal Styles.containerWidth
            , test "has a border radius" <|
                \_ ->
                    cardContainer.borderRadius
                        |> Expect.equal Styles.containerBorderRadius
            , test "has a shadow" <|
                \_ ->
                    cardContainer.shadow
                        |> Expect.equal Styles.containerShadow
            , test "has a hover shadow" <|
                \_ ->
                    cardContainer.hoverShadow
                        |> Expect.equal Styles.containerHoverShadow
            , test "has spacing" <|
                \_ ->
                    cardContainer.spacing
                        |> Expect.equal Styles.containerSpacing
            , test "has left padding" <|
                \_ ->
                    cardContainer.paddingLeft
                        |> Expect.equal Styles.containerPaddingLeft
            ]
        , describe "resource type"
            [ describe "title"
                [ test "has a font size" <|
                    \_ ->
                        cardResourceType.name.size
                            |> Expect.equal Styles.nameSize
                , test "has a font" <|
                    \_ ->
                        cardResourceType.name.font
                            |> Expect.equal Styles.nameFont
                , test "has a font color" <|
                    \_ ->
                        cardResourceType.name.color
                            |> Expect.equal Styles.nameColor
                , test "has vertical padding" <|
                    \_ ->
                        cardResourceType.name.paddingTop
                            |> Expect.equal Styles.namePaddingTop
                , test "has a max width" <|
                    \_ ->
                        cardResourceType.name.maxWidth
                            |> Expect.equal Styles.nameMaxWidth
                ]
            , describe "author"
                [ test "has a font" <|
                    \_ ->
                        cardResourceType.author.font
                            |> Expect.equal Styles.authorFont
                , test "has a font size" <|
                    \_ ->
                        cardResourceType.author.size
                            |> Expect.equal Styles.authorSize
                , test "has a font color" <|
                    \_ ->
                        cardResourceType.author.color
                            |> Expect.equal Styles.authorColor
                , test "has top padding" <|
                    \_ ->
                        cardResourceType.author.paddingTop
                            |> Expect.equal Styles.authorPaddingTop
                ]
            , describe "description"
                [ test "has a font size" <|
                    \_ ->
                        cardResourceType.description.size
                            |> Expect.equal Styles.descriptionSize
                , test "has a font" <|
                    \_ ->
                        cardResourceType.description.font
                            |> Expect.equal Styles.descriptionFont
                , test "has vertical padding" <|
                    \_ ->
                        cardResourceType.description.paddingTop
                            |> Expect.equal Styles.descriptionPaddingTop
                , test "has a font color" <|
                    \_ ->
                        cardResourceType.description.color
                            |> Expect.equal Styles.descriptionColor
                , test "has a max height" <|
                    \_ ->
                        cardResourceType.description.maxHeight
                            |> Expect.equal Styles.descriptionMaxHeight
                , test "has a min height" <|
                    \_ ->
                        cardResourceType.description.minHeight
                            |> Expect.equal Styles.descriptionMinHeight
                , test "has a max width" <|
                    \_ ->
                        cardResourceType.description.maxWidth
                            |> Expect.equal Styles.descriptionMaxWidth
                , test "has spacing for line height" <|
                    \_ ->
                        cardResourceType.description.spacing
                            |> Expect.equal Styles.descriptionSpacing
                ]
            , describe "github"
                [ describe "image"
                    [ test "has a width" <|
                        \_ ->
                            cardResourceType.host.image.width
                                |> Expect.equal Styles.hostImageWidth
                    , test "has a height" <|
                        \_ ->
                            cardResourceType.host.image.height
                                |> Expect.equal Styles.hostImageHeight
                    , test "has a top padding" <|
                        \_ ->
                            cardResourceType.host.image.paddingTop
                                |> Expect.equal Styles.hostImagePaddingTop
                    ]
                , describe "pill"
                    [ test "has a light background color" <|
                        \_ ->
                            cardResourceType.host.pill.lightBackgroundColor
                                |> Expect.equal Styles.hostPillLightBackgroundColor
                    , test "has a dark background color" <|
                        \_ ->
                            cardResourceType.host.pill.darkBackgroundColor
                                |> Expect.equal Styles.hostPillDarkBackgroundColor
                    , test "has a height" <|
                        \_ ->
                            cardResourceType.host.pill.height
                                |> Expect.equal Styles.hostPillHeight
                    , test "has a font face" <|
                        \_ ->
                            cardResourceType.host.pill.font
                                |> Expect.equal Styles.hostPillFont
                    , test "has a font size" <|
                        \_ ->
                            cardResourceType.host.pill.size
                                |> Expect.equal Styles.hostPillFontSize
                    , test "has left padding" <|
                        \_ ->
                            cardResourceType.host.pill.paddingLeft
                                |> Expect.equal Styles.hostPillPaddingLeft
                    , test "has right padding" <|
                        \_ ->
                            cardResourceType.host.pill.paddingRight
                                |> Expect.equal Styles.hostPillPaddingRight
                    , test "has internal spacing" <|
                        \_ ->
                            cardResourceType.host.pill.spacing
                                |> Expect.equal Styles.hostPillSpacing
                    , test "has border radius" <|
                        \_ ->
                            cardResourceType.host.pill.borderRadius
                                |> Expect.equal Styles.hostPillBorderRadius
                    , test "has image height" <|
                        \_ ->
                            cardResourceType.host.pill.imageHeight
                                |> Expect.equal Styles.hostPillImageHeight
                    , test "has image width" <|
                        \_ ->
                            cardResourceType.host.pill.imageWidth
                                |> Expect.equal Styles.hostPillImageWidth
                    ]
                , test "has spacing" <|
                    \_ ->
                        cardResourceType.host.spacing
                            |> Expect.equal Styles.hostSpacing
                ]
            ]
        ]


cardContainer =
    Card.container


cardResourceType =
    Card.resourceType
