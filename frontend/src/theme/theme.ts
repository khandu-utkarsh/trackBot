import { createTheme, ThemeOptions } from '@mui/material/styles';
import { themeConstants } from './constants';
import { Roboto } from 'next/font/google';

const roboto = Roboto({
  weight: ['300', '400', '500', '700'],
  subsets: ['latin'],
  display: 'swap',
});

declare module '@mui/material/styles' {
    interface Theme {
      layout: {
        maxWidth: number | string;
      };
    }
    interface ThemeOptions {
      layout?: {
        maxWidth?: number | string;
      };
    }
}
  

// Shared button styles
const buttonStyles = {
    root: {
        borderRadius: '24px',
        padding: '12px 32px',
        textTransform: 'none',
        fontWeight: 500,
        transition: 'all 0.2s ease-in-out',
    },
    contained: {
        boxShadow: 'none',
        '&:hover': {
            boxShadow: 'none',
            transform: 'translateY(-1px)',
        },
    },
    outlined: {
        borderWidth: '1.5px',
        '&:hover': {
            borderWidth: '1.5px',
            transform: 'translateY(-1px)',
        },
    },
} as const;

// Shared typography settings
const typographySettings = {
    fontFamily: roboto.style.fontFamily,
};

// Shared component overrides
const componentOverrides = {
    MuiButton: {
        styleOverrides: buttonStyles,
    },
};




// Base theme configuration
const baseTheme: ThemeOptions = {
    typography: typographySettings,
    layout: {
        maxWidth: 'md',
    },
};


export const lightTheme = createTheme({
    ...baseTheme,
    palette: {
        mode: 'light',
        primary: {
            main: themeConstants.textColors.primaryLight,
        },
        secondary: {
            main: themeConstants.backgroundColors.secondaryLight,
        },
        background: {
            default: themeConstants.backgroundColors.primaryLight,
            paper: themeConstants.backgroundColors.primaryLight,
        },
        text: {
            primary: themeConstants.textColors.primaryLight,
        },
    },
    components: {
        ...componentOverrides,
        MuiTabs: {
            styleOverrides: {
                root: { 
                    backgroundColor: 'rgba(1, 1, 1, 1)', 
                    borderRadius: '999px', 
                    minHeight: 0 
                },
                indicator: { 
                    display: 'none' 
                },
            },      
        },
        MuiTab: {
            styleOverrides: {
                root: {
                    color: 'white',
                    opacity: 0.7,
                    '&.Mui-selected': {
                        color: 'white',
                        opacity: 1,
                        backgroundColor: 'rgba(255, 255, 255, 0.1)',
                    },
                    minHeight: 40,
                    padding: '8px 16px',
                },
            },
        },
    }
});

export const darkTheme = createTheme({
    ...baseTheme,
    palette: {
        mode: 'dark',
        primary: {
            main: themeConstants.textColors.primaryDark,
        },
        secondary: {
            main: themeConstants.backgroundColors.secondaryDark,
        },
        background: {
            default: themeConstants.backgroundColors.primaryDark,
            paper: themeConstants.backgroundColors.primaryDark,
        },
        text: {
            primary: themeConstants.textColors.primaryDark,
        },
    },
    components: {
        ...componentOverrides,
    },
}); 