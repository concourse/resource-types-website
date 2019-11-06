module BannerTests exposing (..)

import Banner.Banner as Banner exposing (..)
import Test exposing (Test, describe, test)
import Test.Html.Query as Query
import Test.Html.Selector exposing (class, style, tag, text)


suite : Test
suite =
    describe "Banner View"
        [ test "has a container" <|
            \_ ->
                bannerView
                    |> Query.has
                        [ class "banner-container"
                        , style "height" "176px"
                        , style "background-color" "#2A3239"
                        , style "font-family" "Roboto Slab"
                        ]
        , describe "banner title"
            [ test "has the text Concourse Resources" <|
                \_ ->
                    bannerView
                        |> Query.has [ text "Concourse Resources" ]
            , test "has a larger font size" <|
                \_ ->
                    bannerTitle
                        |> Query.has
                            [ style "font-size" "24px" ]
            , test "has a white font color" <|
                \_ ->
                    bannerTitle
                        |> Query.has
                            [ style "color" "#FFFFFF" ]
            , test "center aligns" <|
                \_ ->
                    bannerTitle
                        |> Query.has
                            [ style "display" "grid"
                            , style "align-items" "center"
                            , style "justify-content" "center"
                            ]
            ]
        , describe "banner body"
            [ test "has latin text" <|
                \_ ->
                    bannerBody
                        |> Query.has [ text "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt" ]
            , test "has a white font color" <|
                \_ ->
                    bannerBody
                        |> Query.has
                            [ style "color" "#FFFFFF" ]
            , test "center aligns" <|
                \_ ->
                    bannerBody
                        |> Query.has
                            [ style "display" "grid"
                            , style "align-items" "center"
                            , style "justify-content" "center"
                            ]
            ]
        ]


bannerView =
    Banner.view |> Query.fromHtml


bannerTitle =
    bannerView |> Query.find [ class "banner-title" ]


bannerBody =
    bannerView |> Query.find [ class "banner-body" ]
