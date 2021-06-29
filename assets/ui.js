const gradient1 = 'assets/images/backgrounds/gradient1.png';
const gradient2 = 'assets/images/backgrounds/gradient2.png';
const gradient3 = 'assets/images/backgrounds/gradient3.png';
const gradient4 = 'assets/images/backgrounds/gradient4.png';
const gradient5 = 'assets/images/backgrounds/gradient5.png';
const avatar1 = 'assets/images/avatars/avatar1.png';
const avatar2 = 'assets/images/avatars/avatar2.png';
const avatar3 = 'assets/images/avatars/avatar3.png';
const avatar4 = 'assets/images/avatars/avatar4.png';
const avatar5 = 'assets/images/avatars/avatar5.png';
const avatar6 = 'assets/images/avatars/avatar6.png';
const avatar7 = 'assets/images/avatars/avatar7.png';
const avatar8 = 'assets/images/avatars/avatar8.png';
const avatar9 = 'assets/images/avatars/avatar9.png';
const avatar10 = 'assets/images/avatars/avatar10.png';

window.addEventListener('load', function (event) {
  document.querySelector('.loading').classList.remove('loading--active');
});

// UI functions
const avatarBackgroundRandomizer = () => {
  if (sessionStorage.getItem('avatarBackground')) {
    return (document.querySelector('.avatar').style.backgroundColor = sessionStorage.getItem('avatarBackground'));
  }

  const pastelColors = ['#FFEEC2', '#FFE8E0', '#FFD6E8', '#EDD6FF', '#D6D9FE', '#C5F2E7', '#C0D9DC'];
  const randomColor = pastelColors[Math.floor(Math.random() * pastelColors.length)];

  sessionStorage.setItem('avatarBackground', randomColor);
  document.querySelector('.avatar').style.backgroundColor = randomColor;
};

const gradientRandomizer = () => {
  if (sessionStorage.getItem('gradient')) {
    return (document.querySelector('.backgroundImage').src = sessionStorage.getItem('gradient'));
  }

  const gradients = [gradient1, gradient2, gradient3, gradient4, gradient5];
  const randomGradient = gradients[Math.floor(Math.random() * gradients.length)];

  sessionStorage.setItem('gradient', randomGradient);
  document.querySelector('.backgroundImage').src = randomGradient;
};

const avatarRandomizer = () => {
  if (sessionStorage.getItem('avatarImage')) {
    return (document.querySelector('.avatarImage').src = sessionStorage.getItem('avatarImage'));
  }

  const avatars = [avatar1, avatar2, avatar3, avatar4, avatar5, avatar6, avatar7, avatar8, avatar9, avatar10];
  const randomAvatar = avatars[Math.floor(Math.random() * avatars.length)];

  sessionStorage.setItem('avatarImage', randomAvatar);
  document.querySelector('.avatarImage').src = randomAvatar;
};

gradientRandomizer();
avatarBackgroundRandomizer();
avatarRandomizer();
