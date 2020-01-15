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
                [ test "has a width" <|
                    \_ ->
                        cardResourceType.github.imageWidth
                            |> Expect.equal Styles.githubImageWidth
                , test "has a height" <|
                    \_ ->
                        cardResourceType.github.imageHeight
                            |> Expect.equal Styles.githubImageHeight
                , test "has a top padding" <|
                    \_ ->
                        cardResourceType.github.paddingTop
                            |> Expect.equal Styles.githubImagePaddingTop
                ]
            ]
        ]


cardContainer =
    Card.container


cardResourceType =
    Card.resourceType
