// components/magicui/flickering-grid.tsx
'use client' // Добавьте эту строку в начале файла

import { cn } from '@/lib/utils'
import React, { useEffect, useRef } from 'react'

interface FlickeringGridProps {
	className?: string
	squareSize: number
	gridGap: number
	color: string
	maxOpacity: number
	flickerChance: number
	height: number
	width: number
}

const FlickeringGrid: React.FC<FlickeringGridProps> = ({
	className,
	squareSize,
	gridGap,
	color,
	maxOpacity,
	flickerChance,
	height,
	width,
}) => {
	const canvasRef = useRef<HTMLCanvasElement>(null)

	useEffect(() => {
		const canvas = canvasRef.current
		if (canvas) {
			const ctx = canvas.getContext('2d')
			if (ctx) {
				canvas.width = width
				canvas.height = height

				const drawGrid = () => {
					ctx.clearRect(0, 0, width, height)
					for (let y = 0; y < height; y += squareSize + gridGap) {
						for (let x = 0; x < width; x += squareSize + gridGap) {
							const opacity = Math.random() < flickerChance ? maxOpacity : 0
							ctx.fillStyle = `rgba(${parseInt(
								color.slice(1, 3),
								16
							)}, ${parseInt(color.slice(3, 5), 16)}, ${parseInt(
								color.slice(5, 7),
								16
							)}, ${opacity})`
							ctx.fillRect(x, y, squareSize, squareSize)
						}
					}
				}

				drawGrid()
				const intervalId = setInterval(drawGrid, 100)

				return () => clearInterval(intervalId)
			}
		}
	}, [squareSize, gridGap, color, maxOpacity, flickerChance, height, width])

	return <canvas ref={canvasRef} className={cn('flickering-grid', className)} />
}

export default FlickeringGrid
