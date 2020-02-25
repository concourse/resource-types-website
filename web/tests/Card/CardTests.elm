module Card.CardTests exposing (cardContainer, cardResourceType, suite)

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
                            cardResourceType.github.image.width
                                |> Expect.equal Styles.githubImageWidth
                    , test "has a height" <|
                        \_ ->
                            cardResourceType.github.image.height
                                |> Expect.equal Styles.githubImageHeight
                    , test "has a top padding" <|
                        \_ ->
                            cardResourceType.github.image.paddingTop
                                |> Expect.equal Styles.githubImagePaddingTop
                    ]
                , describe "pill"
                    [ test "has a light background color" <|
                        \_ ->
                            cardResourceType.github.pill.lightBackgroundColor
                                |> Expect.equal Styles.githubPillLightBackgroundColor
                    , test "has a dark background color" <|
                        \_ ->
                            cardResourceType.github.pill.darkBackgroundColor
                                |> Expect.equal Styles.githubPillDarkBackgroundColor
                    , test "has a height" <|
                        \_ ->
                            cardResourceType.github.pill.height
                                |> Expect.equal Styles.githubPillHeight
                    , test "has a font face" <|
                        \_ ->
                            cardResourceType.github.pill.font
                                |> Expect.equal Styles.githubPillFont
                    , test "has a font size" <|
                        \_ ->
                            cardResourceType.github.pill.size
                                |> Expect.equal Styles.githubPillFontSize
                    , test "has left padding" <|
                        \_ ->
                            cardResourceType.github.pill.paddingLeft
                                |> Expect.equal Styles.githubPillPaddingLeft
                    , test "has right padding" <|
                        \_ ->
                            cardResourceType.github.pill.paddingRight
                                |> Expect.equal Styles.githubPillPaddingRight
                    , test "has internal spacing" <|
                        \_ ->
                            cardResourceType.github.pill.spacing
                                |> Expect.equal Styles.githubPillSpacing
                    , test "has border radius" <|
                        \_ ->
                            cardResourceType.github.pill.borderRadius
                                |> Expect.equal Styles.githubPillBorderRadius
                    , test "has image height" <|
                        \_ ->
                            cardResourceType.github.pill.imageHeight
                                |> Expect.equal Styles.githubPillImageHeight
                    , test "has image width" <|
                        \_ ->
                            cardResourceType.github.pill.imageWidth
                                |> Expect.equal Styles.githubPillImageWidth
                    ]
                , test "has spacing" <|
                    \_ ->
                        cardResourceType.github.spacing
                            |> Expect.equal Styles.githubSpacing
                ]
            ]
        ]


cardContainer =
    Card.container


cardResourceType =
    Card.resourceType
