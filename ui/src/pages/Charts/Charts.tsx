import { FC } from 'react';

interface Props {
    canvasRef: React.LegacyRef<HTMLCanvasElement>  // Reference to the HTML canvas element
    width: number                                  // Width of the canvas
    height: number                                 // Height of the canvas
}

const Charts: FC<Props> = ({ canvasRef, height, width }) => {
    return (
        <canvas ref={canvasRef} width={width} height={height}/>
    );
}

export default Charts;