import AnimatedShinyText from '@/components/magicui/animated-shiny-text'
import { BorderBeam } from '@/components/magicui/border-beam'
import { useEffect, useState } from 'react'

const Head = () => {
	const [isClient, setIsClient] = useState(false)

	useEffect(() => {
		setIsClient(true)
	}, [])

	if (!isClient) {
		return null
	}

	const scrollToNextSection = () => {
		const nextSection = document.getElementById('next-section')
		if (nextSection) {
			nextSection.scrollIntoView({ behavior: 'smooth' })
		}
	}

	return (
		<div className='z-10 flex min-h-[50vh] items-center justify-center '>
			<div
				onClick={scrollToNextSection}
				className='relative hover:scale-105  hover:shadow-black/50 transition-all duration-300 flex-col items-center justify-center overflow-hidden rounded-lg  '
			>
				{/* <div className='bg-[#252525] opacity-20 blur-sm absolute inset-10 border-[#6B7280] '></div> */}

				<AnimatedShinyText className='cursor-pointer whitespace-pre-wrap bg-gradient-to-b from-black to-gray-300/80 bg-clip-text text-center text-6xl font-normal leading-none dark:from-white dark:to-slate-900/10 p-5'>
					<div>Hey, I'm Unbas</div>
					<div>I'm a web-developer</div>
					<div>Based in ...</div>
				</AnimatedShinyText>
				<BorderBeam size={300} duration={12} delay={9} className='' />
			</div>
		</div>
	)
}

export default Head
