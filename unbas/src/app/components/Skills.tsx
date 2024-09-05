const Skills = () => {
	const skills = [
		'HTML',
		'CSS',
		'JavaScript',
		'TypeScript',
		'React.js',
		'Next.js',
		'Firebase',
		'Tailwind CSS',
		'Postman',
		'SASS',
		'Git',
		'GitHub',
		'Figma',
		'Magic UI',
		'Next UI',
	]

	return (
		<div className='flex flex-col items-center justify-center h-screen bgCustom p-8'>
			{/* Заголовок */}
			<h1 className='text-9xl md:text-9xl font-bold mb-16 mt-10 text-center'>
				<span className='text-[#252525]'>My Skills</span>
			</h1>

			{/* Список скиллов */}
			<div className='grid grid-cols-2 sm:grid-cols-3 lg:grid-cols-5 gap-10 justify-center'>
				{skills.map((skill, index) => (
					<span
						key={index}
						className='text-2xl md:text-3xl font-normal text-[#252525] text-center'
					>
						{skill}
					</span>
				))}
			</div>
		</div>
	)
}

export default Skills
