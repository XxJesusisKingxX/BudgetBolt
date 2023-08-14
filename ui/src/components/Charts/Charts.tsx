import { FC } from "react";

interface Props {
    canvasRef: React.LegacyRef<HTMLCanvasElement>
    width: number
    height: number
}
const Charts: FC<Props> = ({ canvasRef, height, width }) => {
    return (
        <canvas ref={canvasRef} width={width} height={height}/>
    );
}

export default Charts;