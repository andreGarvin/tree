package allPaths

import (
    "path/filepath"
    "io/ioutil"
    "strings"
    "errors"
    "os"
)

var (
  extString string

  returnedPaths []string
  paths []string
  fs []os.FileInfo

  justDirs bool = false
  justAll bool = true
)


// Returns back all the paths with the given path
func All( path string ) ( []string, error ) {

    err := isAbs(path)
    if err != nil {
        return paths, err
    }

    if err := recursiveTreeDive(path, fs); err != nil {
        return paths, err
    }
    returnedPaths = paths
    paths = []string {}

    return returnedPaths, nil
}

// Returns back the paths that are directories
func Dirs( path string ) ( []string, error ) {

    err := isAbs(path)
    if err != nil {
        return paths, err
    }

    justDirs = true
    justAll = false
    if err := recursiveTreeDive(path, fs); err != nil {
        return paths, err
    }
    returnedPaths = paths
    paths = []string {}

    return returnedPaths, nil
}

// Returns back the paths with a certain file extension
func WithExt( path string, ext string ) ( []string, error ) {

      err := isAbs(path)
      if err != nil {
          return paths, err
      }

      if !strings.Contains(ext, ".") {
          extString = "." + ext
          if err := recursiveTreeDive(path, fs); err != nil {
              return paths, err
          }
          return paths, nil
      }
      if err := recursiveTreeDive(path, fs); err != nil {
          return paths, err
      }
      returnedPaths = paths
      paths = []string {}

      return returnedPaths, nil
}


// recursively gets all files paths in a given directory
func recursiveTreeDive( root string, fs []os.FileInfo ) error {

    for _, f := range fs {

        /*
            gets the file name, concatenate them with the given root,
            then cleans the path slash.
        */
        f := filepath.Clean( filepath.Join( root, f.Name() ) )

        fstat, err := os.Stat( f )
        if err != nil {
            return err
        }

        // checks weather the file is a directory
        if fstat.IsDir() {

            if justDirs {
               paths = append(paths, f )
            }

            newRoot := f

            fs, err := ioutil.ReadDir( newRoot )
            if err != nil {
                return err
            }

            /*
              calls the function agian passing in the 'newRoot'
              and ist directory contents
            */
            if len(fs) != 0 {
                if err := recursiveTreeDive( newRoot, fs ); err != nil {
                    return err
                }
            }
        } else {
            if justAll {

                base := filepath.Base(f)
                if extString != "" && filepath.Ext(base) == extString && filepath.Ext(base) != "" {
                    // appends the all file types name to the 'paths' slice array
                    paths = append(paths, f )
                } else if extString == "" {
                    // appends the all file types name to the 'paths' slice array
                    paths = append(paths, f )
                }
            }
        }
    }
    return nil
}


func isAbs( givenPath string ) error {

    if err := notEmptyString(givenPath); err != nil {
        return err
    }

    absoulte_path, err := filepath.Abs(givenPath)
    if err != nil {
        return err
    }

    if !filepath.IsAbs(absoulte_path) {
        return errors.New("The path provided is not a absolute path")
    }

    localfs, err := ioutil.ReadDir( absoulte_path )
    if err != nil {
        return err
    }

    fs = localfs
    return nil
}

func notEmptyString( str string ) error {

    if len( str ) == 0 {
		    return errors.New("No path provided is a empty string.")
	  }
	  return nil
}

func includes( str string, arr []string ) bool {
    for _, i := range arr {
        if str == i {
            return true
        }
    }
    return false
}

// these are files to gnore and not to cache
// func AllExcept( path string, ignores []string) {
//
//     // checks if there is a cache-ignore file in the current directory
//     if fileExist("./cache-ignore") {
//         // gets the file name/paths from the file
//         data_stream, _ := readFile("./cache-ignore")
//         slice := strings.Split(data_stream, "\n")
//         ignoresSlice = slice[:len( slice ) - 1]
//     } else {
//
//         if len( ignoresSlice ) != 0 {
//             /*
//                 strips the squares brackets from the string
//                 splits the string by semicolons
//             */
//             ignoresSlice = strings.Split( strings.Trim(ignores, "[ ]"), "," )
//         } else {
//             ignoresSlice = []string { ".git", "node_modules" }
//         }
//     }
// }
