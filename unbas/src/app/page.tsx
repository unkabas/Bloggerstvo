'use client'
import FlickeringGridLayout from './FlickeringGridLayout'
import Head from './components/Head'
import Portfolio from './components/Portfolio'
import Skills from './components/Skills'

export default function Home() {
	return (
		<FlickeringGridLayout>
			{/* Добавляем мерцающий фон */}
			<div className=' font-dance h-screen'>
				<section className='relative flex items-center justify-center h-screen'>
					<Head />

					<img
						src='/images/image-1.png'
						alt='Описание изображения'
						className='absolute top-[-800px] left-0 w-full h-auto object-contain opacity-85'
						style={{ height: 'auto', zIndex: -1 }}
					/>
				</section>
				<section
					id='next-section'
					className='relative h-screen bg-custom-image bg-cover bg-center flex justify-center items-center text-black text-5xl hover:opacity-90 transition-opacity duration-300'
				>
					<Skills />
					<img
						src='/images/image-2.png'
						alt='Описание изображения'
						className='absolute top-[-670px] left-0 w-full h-auto object-contain opacity-85'
						style={{ height: 'auto', zIndex: -1 }}
					/>
				</section>
				<section className='bg-custom-image bg-cover bg-center h-screen text-black text-5xl hover:opacity-90 transition-opacity duration-300'>
					<Portfolio />
				</section>
			</div>
		</FlickeringGridLayout>
	)
}
