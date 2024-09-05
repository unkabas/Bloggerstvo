'use client'

import React, { useEffect, useState } from 'react'
import FlickeringGrid from '../components/magicui/flickering-grid'

export default function FlickeringGridLayout({
	children,
}: {
	children: React.ReactNode
}) {
	const [dimensions, setDimensions] = useState({ width: 800, height: 800 })

	useEffect(() => {
		const handleResize = () => {
			setDimensions({
				width: window.innerWidth,
				height: window.innerHeight,
			})
		}

		// Установить начальные размеры
		handleResize()

		// Обновлять размеры при изменении размера окна
		window.addEventListener('resize', handleResize)

		// Удалить слушатель при размонтировании компонента
		return () => {
			window.removeEventListener('resize', handleResize)
		}
	}, [])

	return (
		<div className='relative min-h-screen bg-background overflow-y-auto'>
			{/* Мерцающая сетка */}
			<FlickeringGrid
				className='z-0 fixed top-0 left-0 w-full h-full pointer-events-none'
				squareSize={1.5}
				gridGap={20}
				color='#6B7280'
				maxOpacity={1}
				flickerChance={0.1}
				height={dimensions.height}
				width={dimensions.width}
			/>
			<div className='relative z-10'>{children}</div>
		</div>
	)
}
