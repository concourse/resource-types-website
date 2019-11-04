module MainTests exposing (..)

import Main exposing (..)
import Test exposing (Test, describe, test)
import Test.Html.Query as Query
import Test.Html.Selector exposing (class, style, tag, text)


suite : Test
suite =
    describe "Banner"
        [ test "should have the text Concourse Resources" <|
            \_ ->
                Main.view {}
                    |> Query.fromHtml
                    |> Query.has [ text "Concourse Resources" ]
        , test "should have a container" <|
            \_ ->
                Main.view {}
                    |> Query.fromHtml
                    |> Query.has [ class "banner-title" ]
        , test "container should have a height of 176px" <|
            \_ ->
                Main.view {}
                    |> Query.fromHtml
                    |> Query.find [ class "banner-title" ]
                    |> Query.has [ style "height" "176px" ]
        , test "container should be center aligned" <|
            \_ ->
                Main.view {}
                    |> Query.fromHtml
                    |> Query.find [ class "banner-title" ]
                    |> Query.has [ style "text-align" "center" ]
        ]
