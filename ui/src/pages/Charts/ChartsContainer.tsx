import { useEffect, useRef } from 'react';
import Charts from './Charts';

const ChartsContainer = () => {
    // Create a ref to hold the canvas element
    const ref = useRef<HTMLCanvasElement | null>(null);

    // Generate a random number within a specified range
    const random = (max: number, min: number) => {
        return Math.random() * (max - min) + min;
    };

    // useEffect to draw the chart on the canvas
    useEffect(() => {
        // Access the canvas element from the ref
        const canvas = ref.current;

        if (canvas) {
            // Get the 2D rendering context of the canvas
            const context = canvas.getContext("2d");

            if (context) {
                // Adjust the coordinate system
                context.scale(1, -1);
                context.translate(0, -canvas.height);

                // Configure drawing settings
                context.strokeStyle = "black";
                context.lineWidth = 25;
                context.lineCap = "round";

                // Start drawing path
                context.beginPath();
                context.moveTo(0, 0);

                // Create a line on the chart
                let i = 200;
                while (i < 10450) {
                    const rand = random(0, 1000);
                    context.lineTo(i, rand);
                    i += 200;
                }

                // Draw the line
                context.stroke();
            }
        }
    }, []);

    return (
        // Render the Charts component with canvas and dimensions
        <Charts
            canvasRef={ref}     // Reference to the canvas element
            width={10450}       // Width of the canvas
            height={1100}       // Height of the canvas
        />
    );
}

export default ChartsContainer;
