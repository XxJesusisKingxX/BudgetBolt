import { useEffect, useRef } from "react";
import Charts from "./Charts";

const ChartsContainer = () => {
    const ref = useRef<HTMLCanvasElement | null>(null);
    const random = (max: number, min: number) => {
        return Math.random() * (max - min) + min;
    }
    useEffect(() => {
        const canvas = ref.current;
        if (canvas) {
            const context = canvas.getContext("2d");
            if (context) {
                context.scale(1, -1);
                context.translate(0, -canvas.height);
                context.strokeStyle = "black";
                context.lineWidth = 25;
                context.lineCap = "round";
                context.beginPath();
                context.moveTo(0, 0);
                // Create line
                let i = 200
                while (i < 10450) {
                    const rand = random(0,1000)
                    context.lineTo(i, rand)
                    i += 200
                }
                context.stroke();
            }
        }
    }, []);

    return (
        <Charts
            canvasRef={ref}
            width={10450}
            height={1100}
        />
    );
}

export default ChartsContainer