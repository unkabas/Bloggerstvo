import { BorderBeam } from '@/components/magicui/border-beam'
import { useState } from 'react'

export default function Portfolio() {
	const [hoveredIndex, setHoveredIndex] = useState<number | null>(null)

	const list = [
		{
			title: 'Neyasot',
			img: '/images/neyasot.png', // Make sure the image exists in the public/images folder
		},
		// Add more projects here if needed
	]

	return (
		<div className='p-8 flex justify-center min-h-screen'>
			{/* Portfolio Section */}
			<div className='w-1/2 flex items-center'>
				{/* Content aligned to the center */}
				<div className='flex flex-col items-center ml-24'>
					{list.map((item, index) => (
						<div
							key={index}
							onMouseEnter={() => setHoveredIndex(index)} // Set the hovered index on mouse enter
							onMouseLeave={() => setHoveredIndex(null)} // Reset the hovered index on mouse leave
							className='flex justify-center items-center w-52 p-4 mb-4 transition-all duration-300 transform hover:bg-gray-200 rounded-full hover:shadow-md hover:ml-4 cursor-pointer'
						>
							<div className='flex justify-center items-center'>
								<p className='text-2xl font-normal'>{item.title}</p>
							</div>
						</div>
					))}
				</div>
			</div>

			{/* Display image on hover */}
			<div className='w-full flex justify-center items-center mr-40'>
				{hoveredIndex !== null && (
					<div className='relative'>
						{/* Image */}
						<img
							src={list[hoveredIndex].img}
							alt={list[hoveredIndex].title}
							className='object-contain rounded-lg shadow-lg transition-opacity duration-300'
						/>
						{/* BorderBeam positioned around the image */}
						<BorderBeam
							size={600} // Make sure size is slightly larger than image
							duration={12}
							delay={9}
							className='absolute inset-0 rounded-lg'
						/>
					</div>
				)}
			</div>
		</div>
	)
}
