// return Hue
export const yugiohCardColor = (attribute: string, cardtype: string[]): string => {
  if (attribute.includes("魔法")) {
    return "#3A633A"
  }
  if (attribute.includes("罠")) {
    return "#633A55"
  }

  if (cardtype.includes("融合")) {
    return "#533A63"
  }
  if (cardtype.includes("儀式")) {
    return "#313054"
  }
  if (cardtype.includes("エクシーズ")) {
    return "#0e0e0e"
  }
  if (cardtype.includes("シンクロ")) {
    return "#b2afb4"
  }
  if (cardtype.includes("通常")) {
    return "#8C795D"
  }
  if (cardtype.includes("効果")) {
    return "#634636"
  }
  return "#ffffff"
}

