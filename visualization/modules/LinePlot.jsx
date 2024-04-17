import * as d3 from "d3";
import { useEffect } from "react";

export default function LinePlot({
  data,
  width = 640,
  height = 400,
  marginTop = 20,
  marginRight = 20,
  marginBottom = 20,
  marginLeft = 20
}) {
  useEffect(() => {
    const data = [12, 2, 3, 4, 50, 1]
    demo(data, ".demo")
  }, [])

  return (
    <svg class="demo"></svg>
  );
}

/**
 * 
 * @param {any[]} data 
 * @param {string} selector 
 */
function demo(data, selector) {
  const svg = d3.select(selector)
    .attr("width", 700)
    .attr("height", 300);

  const scaleLiner = d3.scaleLinear()


  svg.selectAll("rect")
    .data(data)
    .enter()
    .append("rect")
    .attr("x", (d, i) => i * 70)
    .attr("y", (d, i) => 200 - 10 * d)
    .attr("width", 65)
    .attr("height", (d, i) => d * 10)
    .attr("fill", "green")

  svg.selectAll("text")
    .data(data)
    .enter()
    .append("text")
    .text((d) => d)
    .attr("x", (d, i) => i * 70 + 30)
    .attr("y", (d, i) => 200 - 10 * d - 3)
    .attr("fill", "white")
    .attr("text-anchor", "middle")
}
