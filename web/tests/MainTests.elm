module MainTests exposing (..)

import Main exposing (view)
import Test exposing (Test, describe, test)
import Test.Html.Query as Query
import Test.Html.Selector exposing (class, style)


suite : Test
suite =
    describe "Main"
        [ test "has a set width" <|
            \_ ->
                Main.view
                    |> Query.fromHtml
                    |> Query.has
                        [ style "width" "1024px" ]
        , test "is centered" <|
            \_ ->
                Main.view
                    |> Query.fromHtml
                    |> Query.has
                        [ style "margin" "0 auto"
                        , style "width" "1024px"
                        ]
        ]
