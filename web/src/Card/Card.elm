module Card.Card exposing (Author, Card, Container, Description, Github, Name, ResourceType, card, container, resourceType)

import Card.Styles as Styles
import Common.Common exposing (RGB, Shadow)


type alias Card =
    { container : Container
    , resourceType : ResourceType
    }


type alias Container =
    { height : Int
    , width : Int
    , borderRadius : Int
    , shadow : Shadow
    , hoverShadow : Shadow
    , spacing : Int
    , paddingLeft : Int
    }


type alias ResourceType =
    { name : Name
    , author : Author
    , description : Description
    , github : Github
    }


type alias Name =
    { size : Int
    , font : String
    , paddingTop : Int
    , maxWidth : Int
    , color : RGB
    }


type alias Author =
    { font : String
    , size : Int
    , color : RGB
    , paddingTop : Int
    }


type alias Description =
    { size : Int
    , font : String
    , paddingTop : Int
    , color : RGB
    , maxHeight : Int
    , minHeight : Int
    , maxWidth : Int
    , spacing : Int
    }


type alias Github =
    { image : GithubImage
    , pill : GithubPill
    , spacing : Int
    }


type alias GithubImage =
    { height : Int
    , width : Int
    , paddingTop : Int
    }


type alias GithubPill =
    { lightBackgroundColor : RGB
    , darkBackgroundColor : RGB
    , height : Int
    , size : Int
    , font : String
    , borderRadius : Int
    , imageHeight : Int
    , imageWidth : Int
    , paddingLeft : Int
    , paddingRight : Int
    , spacing : Int
    }


card : Card
card =
    { container = container
    , resourceType = resourceType
    }


container : Container
container =
    { height = Styles.containerHeight
    , width = Styles.containerWidth
    , borderRadius = Styles.containerBorderRadius
    , shadow = Styles.containerShadow
    , hoverShadow = Styles.containerHoverShadow
    , spacing = Styles.containerSpacing
    , paddingLeft = Styles.containerPaddingLeft
    }


resourceType : ResourceType
resourceType =
    { name = name
    , author = author
    , description = description
    , github = github
    }


name : Name
name =
    { size = Styles.nameSize
    , font = Styles.nameFont
    , paddingTop = Styles.namePaddingTop
    , maxWidth = Styles.nameMaxWidth
    , color = Styles.nameColor
    }


author : Author
author =
    { font = Styles.authorFont
    , size = Styles.authorSize
    , color = Styles.authorColor
    , paddingTop = Styles.authorPaddingTop
    }


description : Description
description =
    { size = Styles.descriptionSize
    , font = Styles.descriptionFont
    , paddingTop = Styles.descriptionPaddingTop
    , color = Styles.descriptionColor
    , maxHeight = Styles.descriptionMaxHeight
    , minHeight = Styles.descriptionMinHeight
    , maxWidth = Styles.descriptionMaxWidth
    , spacing = Styles.descriptionSpacing
    }


github : Github
github =
    { image = image
    , pill = pill
    , spacing = Styles.githubSpacing
    }


image : GithubImage
image =
    { height = Styles.githubImageHeight
    , width = Styles.githubImageWidth
    , paddingTop = Styles.githubImagePaddingTop
    }


pill : GithubPill
pill =
    { lightBackgroundColor = Styles.githubPillLightBackgroundColor
    , darkBackgroundColor = Styles.githubPillDarkBackgroundColor
    , height = Styles.githubPillHeight
    , font = Styles.githubPillFont
    , size = Styles.githubPillFontSize
    , borderRadius = Styles.githubPillBorderRadius
    , imageHeight = Styles.githubPillImageHeight
    , imageWidth = Styles.githubPillImageWidth
    , paddingLeft = Styles.githubPillPaddingLeft
    , paddingRight = Styles.githubPillPaddingRight
    , spacing = Styles.githubPillSpacing
    }
